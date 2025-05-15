package dto

import (
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"
)

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
		Status     string `json:"status"`
	}

	ClassSettingListResponse struct {
		ClassSetting []ClassSettingResponse `json:"class_settings"`
		Meta         meta.Meta              `json:"meta"`
	}

	ClassSettingList struct {
		ClassSetting []entity.ClassSettings `json:"class_settings"`
		Meta         meta.Meta              `json:"meta"`
	}
)
