package auth

import (
	// "errors"
	"rpba-app/pkg/models"
	"strings"
	"sync"
)

var (
	AuthUser *models.User
	mu sync.RWMutex
)


func SetAuthUser(user *models.User) {
	mu.Lock()
	defer mu.Unlock()
	
	AuthUser = user
}

func GetAuthUser() *models.User {
	mu.RLock()
	defer mu.RUnlock()
	return AuthUser
}

func UserHasRole(role string) (bool) {
	user := GetAuthUser()
	if user == nil {
		return false
	}
	
	for _, r := range user.Roles {
		if strings.EqualFold(r.Name, role) {
			return true
		}
	}
	return false
}

func UserHasPermission(permission string) (bool) {
	user := GetAuthUser()
	if user == nil {
		return false
	}

	// Loop through the user's roles and check for the desired permission
	for _, r := range user.Roles {
		for _, p := range r.Permissions {
			if strings.EqualFold(p.Name, permission) {
				return true
			}
		}
	}
	return false
}