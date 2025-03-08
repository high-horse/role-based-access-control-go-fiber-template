package auth

import (
	"log"
	"net/http"
	"rbac/db/pool"
	"rbac/helpers"
	"rbac/pkg/models"
	"rbac/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Login(c *fiber.Ctx) error {
	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&loginRequest); err != nil {
		helpers.ResponseError(c, http.StatusBadRequest, "Invalid Request")
	}

	var user models.User
	result := pool.DB.Where("username = ?", loginRequest.Username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// User not found, return error
			return helpers.ResponseError(c, http.StatusUnauthorized, "Invalid Username or Password")
		}
		// Some other error occurred
		return helpers.ResponseError(c, http.StatusInternalServerError, "Internal Server Error")
	}

	token, err := utils.GenerateToken(loginRequest.Password, []string{}, []string{})
	if err != nil {
		helpers.ResponseError(c, http.StatusInternalServerError, "Internal Server error")
	}

	return helpers.ResponseJson(c, http.StatusOK, struct {
		Status bool   `json:"status"`
		Token  string `json:"token"`
	}{
		Status: true,
		Token:  token,
	})
}

func Register(c *fiber.Ctx) error {
	var registerForm struct {
		Username string `json:"username" validate:"required,min=5,max=20"`
		Name     string `json:"name" validate:"required,min=5,max=50"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
		Role     int    `json:"role" validate:"required"` // Role must be either "admin" or "user"
	}
	// Parse incoming JSON into the registerForm struct
	if err := c.BodyParser(&registerForm); err != nil {
		helpers.ResponseError(c, http.StatusBadRequest, "Invalid Request")
		return err
	}

	utils.LogToFile(registerForm)
	validationErrors, err := utils.ValidateRequest(registerForm)
	if err != nil {
		return helpers.ResponseError(c, 500, "Internal Server Error")
	}
	// If there are validation errors, use the helper to handle them
	if len(validationErrors) > 0 {
		// Respond with a 422 status code for validation errors.
		return helpers.ResponseValidationError(c, validationErrors)
	}

	var existingUser models.User
	result := pool.DB.Where("username = ?", registerForm.Username).First(&existingUser)
	if result.Error == nil {
		return helpers.ResponseError(c, 409, "Username already exists")
	}

	hashedPw, err := utils.HashPassowrd(registerForm.Password)
	log.Println("hashed pass ", hashedPw)
	if result.Error == nil {
		return helpers.ResponseError(c, 500, err.Error())
	}
	newUser := models.User{
		Username: registerForm.Username,
		Name:     registerForm.Name,
		Email:    registerForm.Email,
		Password: hashedPw,
		RoleID:   uint(registerForm.Role),                     
	}

	if err := pool.DB.Create(&newUser).Error; err != nil {
		return helpers.ResponseError(c, 500, err.Error())
	}

	// Return success response
	return helpers.ResponseJson(c, 200, struct {
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}{
		Status:  true,
		Message: "Successfully created user",
	})
}
