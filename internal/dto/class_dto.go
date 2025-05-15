package dto

type (
	CreateClassRequest struct {
		Lecturer  string `json:"lecturer" binding:"required"`
		CourseID  string `json:"course_id" binding:"required"`
		Day       string `json:"day" binding:"required"` // changed from time.Time to string for day
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Classroom string `json:"classroom" binding:"required"`
	}

	UpdateClassRequest struct {
		Lecturer  string `json:"lecturer" binding:"required"`
		CourseID  string `json:"course_id" binding:"required"`
		Day       string `json:"day" binding:"required"` // changed from time.Time to string for day
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Classroom string `json:"classroom" binding:"required"`
	}

	ClassResponse struct {
		ID        string `json:"id"`
		Lecturer  string `json:"lecturer"`
		CourseID  string `json:"course_id"`
		Day       string `json:"day"`
		StartTime string `json:"start_time"`
		EndTime   string `json:"end_time"`
		Classroom string `json:"classroom"`
	}
)
