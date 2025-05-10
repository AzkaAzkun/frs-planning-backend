package dto

type CreateCourseRequest struct {
	Name           string `json:"name" binding:"required"`
	ClassSettingID int64  `json:"class_setting_id" binding:"required"`
}

type UpdateCourseRequest struct {
	Name           string `json:"name" binding:"required"`
	ClassSettingID int64  `json:"class_setting_id" binding:"required"`
}

type CourseResponse struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ClassSettingID int64  `json:"class_setting_id"`
}
