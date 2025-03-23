package models

import (
	"gorm.io/gorm"
)

// Exercise model
type Exercise struct {
	gorm.Model
	Name        string    `gorm:"size:100;not null" json:"name"`
	Description string    `json:"description"`
	CategoryID  uint      `json:"category_id"`
	Category    Category  `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	MuscleGroup string    `gorm:"size:50" json:"muscle_group"`
}
