package models

import (
	"time"

	"gorm.io/gorm"
)

type WorkoutPlan struct {
	gorm.Model
	UserID      uint              `json:"user_id"`
	Title       string            `gorm:"size:100;not null" json:"title"`
	Description string            `json:"description"`
	Exercises   []WorkoutExercise `gorm:"foreignKey:WorkoutPlanID" json:"exercises,omitempty"`
	Sessions    []WorkoutSession  `gorm:"foreignKey:WorkoutPlanID" json:"sessions,omitempty"`
}

// WorkoutExercise junction model
type WorkoutExercise struct {
	gorm.Model
	WorkoutPlanID uint     `json:"workout_plan_id"`
	ExerciseID    uint     `json:"exercise_id"`
	Exercise      Exercise `gorm:"foreignKey:ExerciseID" json:"exercise"`
	Sets          int      `gorm:"default:3" json:"sets"`
	Reps          int      `gorm:"default:10" json:"reps"`
	Weight        float64  `json:"weight,omitempty"`
	Duration      int      `json:"duration,omitempty"`
	Notes         string   `json:"notes,omitempty"`
	OrderPosition int      `gorm:"default:0" json:"order_position"`
}

// WorkoutSession model
type WorkoutSession struct {
	gorm.Model
	WorkoutPlanID uint            `json:"workout_plan_id"`
	ScheduledTime time.Time       `json:"scheduled_time"`
	CompletedTime *time.Time      `json:"completed_time,omitempty"`
	Status        string          `gorm:"size:20;default:'scheduled'" json:"status"` // scheduled, completed, skipped
	Notes         string          `json:"notes,omitempty"`
	Results       []WorkoutResult `gorm:"foreignKey:WorkoutSessionID" json:"results,omitempty"`
}

// WorkoutResult model
type WorkoutResult struct {
	gorm.Model
	WorkoutSessionID uint     `json:"workout_session_id"`
	ExerciseID       uint     `json:"exercise_id"`
	Exercise         Exercise `gorm:"foreignKey:ExerciseID" json:"exercise"`
	SetsCompleted    int      `gorm:"default:0" json:"sets_completed"`
	RepsCompleted    int      `gorm:"default:0" json:"reps_completed"`
	WeightUsed       float64  `json:"weight_used,omitempty"`
	DurationSeconds  int      `json:"duration_seconds,omitempty"`
	Notes            string   `json:"notes,omitempty"`
}
