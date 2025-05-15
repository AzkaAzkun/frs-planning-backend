package dto

import (
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"
)

type (
	PlanCreateRequest struct {
		Name        string `json:"name" binding:"required"`
		WorkspaceID string `json:"workspace_id" binding:"required"`
	}

	PlanResponse struct {
		ID          string `json:"id"`
		WorkspaceID string `json:"workspace_id"`
		Name        string `json:"name"`
	}

	GetAllPlan struct {
		Plans []entity.Plan `json:"plans"`
		Meta  meta.Meta     `json:"meta"`
	}

	GetAllPlanResponse struct {
		Plans []PlanResponse `json:"plans"`
		Meta  meta.Meta      `json:"meta"`
	}
)
