package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null"`
	Slug        string         `gorm:"not null"`
	Description string
	CreatedBy   uint           `gorm:"index"` // Foreign key to User (creator)
	CreatedAt   time.Time      `gorm:"default:now()"`
	UpdatedAt   time.Time      `gorm:"default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	// Creator of the role
	// Creator     User          `gorm:"foreignKey:CreatedBy;constraint:OnDelete:CASCADE;"`

	// Many-to-many relationship with Permissions through the 'role_has_permissions' table
	Permissions []Permission `gorm:"many2many:role_has_permissions;"`
}
