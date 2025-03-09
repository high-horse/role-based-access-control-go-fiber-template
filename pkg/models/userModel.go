package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	Status    bool   `gorm:"default:true"`
	RoleID    uint   `gorm:"not null"` // Foreign key to Role
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Relationship to Role through the User's RoleID
	Role Role `gorm:"foreignKey:RoleID"`
	
	// Permission []Permission ``

	// Many-to-many relationship with Role through the 'role_has_permissions' table
	// Roles []Role `gorm:"many2many:role_has_permissions;"`
}



