package models

import (
	// "gorm.io/gorm"
)

// RoleHasPermissions is a join table for roles and permissions
type RoleHasPermissions struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint `gorm:"index"`
	PermissionID uint `gorm:"index"`

	Role       Role       `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE;"`
	Permission Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE;"`
}
