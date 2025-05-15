package dto

import "time"

type (
	CreateClassRequest struct {
		Lecturer      string `json:"lecturer" binding:"required"`
		CourseID      string `json:"course_id" binding:"required"`
		ClassSchedule string `json:"class_schedule" binding:"required"` // changed from time.Time to string for time only
		Priority      int    `json:"priority" binding:"required"`
		Classroom     string `json:"classroom" binding:"required"`
	}

	UpdateClassRequest struct {
		Lecturer      string `json:"lecturer" binding:"required"`
		CourseID      string `json:"course_id" binding:"required"`
		ClassSchedule string `json:"class_schedule" binding:"required"` // changed from time.Time to string for time only
		Priority      int    `json:"priority" binding:"required"`
		Classroom     string `json:"classroom" binding:"required"`
	}

	ClassResponse struct {
		ID            string    `json:"id"`
		Lecturer      string    `json:"lecturer"`
		CourseID      string    `json:"course_id"`
		ClassSchedule time.Time `json:"class_schedule"`
		Priority      int       `json:"priority"`
		Classroom     string    `json:"classroom"`
	}
)
