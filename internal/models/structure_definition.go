package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type StructureDefinition struct {
	ID          string         `gorm:"primaryKey"`
	URL         string         `gorm:"uniqueIndex"`
	Name        string
	Title       string
	Status      string
	Kind        string
	Type        string
	Content     datatypes.JSON `gorm:"type:jsonb"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
