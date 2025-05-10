package dto

import "time"

type CreateClassRequest struct {
	Lecturer      string    `json:"lecturer" binding:"required"`
	CourseID      string    `json:"course_id" binding:"required"`
	ClassSchedule time.Time `json:"class_schedule" binding:"required"`
	Priority      int       `json:"priority" binding:"required"`
}

type UpdateClassRequest struct {
	Lecturer      string    `json:"lecturer" binding:"required"`
	CourseID      string    `json:"course_id" binding:"required"`
	ClassSchedule time.Time `json:"class_schedule" binding:"required"`
	Priority      int       `json:"priority" binding:"required"`
}

type ClassResponse struct {
	ID            int64     `json:"id"`
	Lecturer      string    `json:"lecturer"`
	CourseID      string    `json:"course_id"`
	ClassSchedule time.Time `json:"class_schedule"`
	Priority      int       `json:"priority"`
}
