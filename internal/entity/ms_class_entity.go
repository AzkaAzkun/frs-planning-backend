package entity

import (
	"github.com/google/uuid"
)

type Class struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Lecturer  string    `json:"lecturer"`
	CourseID  uuid.UUID `json:"course_id"`
	Name      string    `json:"name"`
	Day       string    `json:"day"`
	StartTime string    `json:"start_time"`
	EndTime   string    `json:"end_time"`
	Classroom string    `json:"classroom"`

	Course *Course `json:"course" gorm:"foreignKey:CourseID;references:ID"`
	Timestamp
}

func (Class) TableName() string {
	return "classes"
}
