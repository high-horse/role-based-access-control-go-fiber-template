package models

import (
	"time"

	"gorm.io/gorm"
)

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Slug        string `gorm:"not null"`
	Description string
	CreatedBy   uint           `gorm:"index"` // Foreign key to User (creator)
	CreatedAt   time.Time      `gorm:"default:now()"`
	UpdatedAt   time.Time      `gorm:"default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Creator of the permission
	// Creator     User          `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE;"`
	
	// Many-to-many relationship with Roles through the 'role_has_permissions' table
	// Roles       []Role        `gorm:"many2many:role_has_permissions;"`
}
