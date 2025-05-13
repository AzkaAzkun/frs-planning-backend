package dto

type (
	CreateWorkspaceRequest struct {
		UserID string `json:"user_id" binding:"required"`
		Name   string `json:"name" binding:"required"`
	}

	UpdateWorkspaceRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	DeleteWorkspaceRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	WorkspaceResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
