package entity

import "time"

type Class struct {
	ID            int64     `json:"id" gorm:"primaryKey"`
	Lecturer      string    `json:"lecturer"`
	CourseID      string    `json:"course_id"`
	ClassSchedule time.Time `json:"class_schedule"`
	Priority      int       `json:"priority"`
}

func (Class) TableName() string {
	return "classes"
}
