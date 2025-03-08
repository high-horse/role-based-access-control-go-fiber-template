package models

import (
	"time"

	"gorm.io/gorm"
)

// Posts for the User
type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Description string `gorm:"size:1000"`
	Body        string
	UserID      uint           `gorm:"index"` // Foreign key to User
	CreatedAt   time.Time      `gorm:"default:now()"`
	UpdatedAt   time.Time      `gorm:"default:now()"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`

	User User `gorm:"constraint:OnDelete:CASCADE;"`
}
