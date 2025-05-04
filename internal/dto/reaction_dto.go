package dto

type (
	ReactionRequest struct {
		ReviewId string `json:"review_id" binding:"required,uuid"`
		Status   string `json:"status" binding:"required"`
	}

	ReactionUpdate struct {
		Status string `json:"status" binding:"required"`
	}
)
