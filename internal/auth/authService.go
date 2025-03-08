package auth

import (
	"errors"
	"fmt"
	"rbac/db/pool"
	"rbac/pkg/models"
	"rbac/pkg/utils"

	"gorm.io/gorm"
)

func AuthenticateUser(username, password string) (string, error) {
	var user models.User

	if err := pool.DB.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("record not found")
		}

		return "", err
	}

	if result := utils.CheckPasswordHash(password, user.Password); !result {
		return "", errors.New("invalid password")
	}

	var roles []string
	var permissions []string

	roles = append(roles, user.Role.Name)

	// Add the permission names associated with the user's role
	for _, permission := range user.Role.Permissions {
		permissions = append(permissions, permission.Name)
	}
	
	token, err := utils.GenerateToken(user.Username, roles, permissions)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil

}


