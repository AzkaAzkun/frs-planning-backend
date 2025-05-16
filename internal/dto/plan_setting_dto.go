package dto

type (
	PlanSettingCreateRequest struct {
		PlanID  string `json:"plan_id" binding:"required"`
		ClassID string `json:"class_id" binding:"required"`
	}

	PlanSettingResponse struct {
		ID      string        `json:"id"`
		ClassID string        `json:"class_id"`
		Class   ClassResponse `json:"class"`
	}
)
