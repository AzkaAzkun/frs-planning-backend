package dto

type (
	CreateCourseRequest struct {
		Name           string `json:"name" binding:"required"`
		ClassSettingID string `json:"class_setting_id" binding:"required,uuid"`
	}

	UpdateCourseRequest struct {
		Name           string `json:"name" binding:"required"`
		ClassSettingID string `json:"class_setting_id" binding:"required"`
	}

	CourseResponse struct {
		ID             string `json:"id"`
		Name           string `json:"name"`
		ClassSettingID string `json:"class_setting_id"`
	}
)
