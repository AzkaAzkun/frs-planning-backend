package dto

type (
	CreateClassSettingRequest struct {
		Name       string `json:"name" binding:"required"`
		Permission string `json:"permission" binding:"required"`
	}

	CloneClassSettingRequest struct {
		ClassSettingId string `json:"class_setting_id" binding:"required"`
	}

	ClassSettingResponse struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		User_id    string `json:"user_id"`
		Permission string `json:"permission"`
	}
)
