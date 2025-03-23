package models

import "gorm.io/gorm"

type user struct {
	gorm.Model
	UserID       uint          `gorm:"primaryKey" json:"user_id"`
	Username     string        `gorm:"size:50;not null;unique" json:"username"`
	PasswordHash string        `gorm:"not null" json:"-"`
	Email        string        `gorm:"size:100;not null;unique" json:"email"`
	WorkoutPlans []WorkoutPlan `gorm:"foreignKey:UserID" json:"workout_plans,omitempty"`
}
