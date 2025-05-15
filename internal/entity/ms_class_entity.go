package entity

import (
	"time"

	"github.com/google/uuid"
)

type Class struct {
	ID            uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Lecturer      string    `json:"lecturer"`
	CourseID      uuid.UUID `json:"course_id"`
	ClassSchedule time.Time `json:"class_schedule"`
	Priority      int       `json:"priority"`
	Classroom     string    `json:"classroom"`

	Timestamp
}

func (Class) TableName() string {
	return "classes"
}
