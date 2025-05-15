package dto

import (
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"
)

type (
	CreateWorkspaceRequest struct {
		Name string `json:"name" binding:"required"`
	}

	UpdateWorkspaceRequest struct {
		ID   string `json:"id" binding:"required"`
		Name string `json:"name" binding:"required"`
	}

	DeleteWorkspaceRequest struct {
		ID string `json:"id" binding:"required"`
	}

	GetAllWorkspaceResponse struct {
		Workspaces []WorkspaceResponse `json:"workspaces"`
		Meta       meta.Meta           `json:"meta"`
	}

	GetAllWorkspace struct {
		Workspaces []entity.Workspace `json:"workspaces"`
		Meta       meta.Meta          `json:"meta"`
	}

	WorkspaceResponse struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
