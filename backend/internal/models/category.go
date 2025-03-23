package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"size:50;not null;unique" json:"name"`
	Description string    `json:"description"`
	Exercises   []Exercise `gorm:"foreignKey:CategoryID" json:"exercises,omitempty"`

}