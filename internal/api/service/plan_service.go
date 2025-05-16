package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	PlanService interface {
		CreatePlan(ctx context.Context, req dto.PlanCreateRequest) (dto.PlanResponse, error)
		GetAllPlans(ctx context.Context, workspaceId string, metareq meta.Meta) (dto.GetAllPlanResponse, error)
		GetPlanDetail(ctx context.Context, planId string) (dto.PlanDetailResponse, error)
		Update(ctx context.Context, planId string, req dto.PlanUpdateRequest) error
		Delete(ctx context.Context, planId string) error
	}

	planService struct {
		planRepo      repository.PlanRepository
		workspaceRepo repository.WorkspaceRepository
		db            *gorm.DB
	}
)

func NewPlanService(planRepo repository.PlanRepository,
	workspaceRepo repository.WorkspaceRepository,
	db *gorm.DB) PlanService {
	return &planService{
		planRepo:      planRepo,
		workspaceRepo: workspaceRepo,
		db:            db,
	}
}

func (s *planService) CreatePlan(ctx context.Context, req dto.PlanCreateRequest) (dto.PlanResponse, error) {
	workspace, err := s.workspaceRepo.Find(ctx, nil, uuid.MustParse(req.WorkspaceID))
	if err != nil {
		return dto.PlanResponse{}, err
	}

	plan, err := s.planRepo.Create(ctx, nil, entity.Plan{
		WorkspaceID: workspace.ID,
		Name:        req.Name,
	})
	if err != nil {
		return dto.PlanResponse{}, err
	}

	return dto.PlanResponse{
		ID:          plan.ID.String(),
		WorkspaceID: plan.WorkspaceID.String(),
		Name:        plan.Name,
	}, nil
}

func (s *planService) GetAllPlans(ctx context.Context, workspaceId string, metareq meta.Meta) (dto.GetAllPlanResponse, error) {
	plans, err := s.planRepo.FindAll(ctx, nil, workspaceId, metareq)
	if err != nil {
		return dto.GetAllPlanResponse{}, err
	}

	var planResponses []dto.PlanResponse
	for _, plan := range plans.Plans {
		planResponses = append(planResponses, dto.PlanResponse{
			ID:          plan.ID.String(),
			WorkspaceID: plan.WorkspaceID.String(),
			Name:        plan.Name,
		})
	}

	return dto.GetAllPlanResponse{
		Plans: planResponses,
		Meta:  plans.Meta,
	}, nil
}

func (s *planService) GetPlanDetail(ctx context.Context, planId string) (dto.PlanDetailResponse, error) {
	plan, err := s.planRepo.FindByID(ctx, nil, planId)
	if err != nil {
		return dto.PlanDetailResponse{}, err
	}

	var planResponse []dto.PlanSettingResponse
	for _, planSetting := range plan.PlanSettings {
		planResponse = append(planResponse, dto.PlanSettingResponse{
			ID:      planSetting.ID.String(),
			ClassID: planSetting.Class.ID.String(),
			Class: dto.ClassResponse{
				ID:        planSetting.Class.ID.String(),
				Lecturer:  planSetting.Class.Lecturer,
				Name:      planSetting.Class.Name,
				Classroom: planSetting.Class.Classroom,
				Day:       planSetting.Class.Day,
				StartTime: planSetting.Class.StartTime,
				EndTime:   planSetting.Class.EndTime,
			},
		})
	}

	return dto.PlanDetailResponse{
		ID:           plan.ID.String(),
		WorkspaceID:  plan.WorkspaceID.String(),
		Name:         plan.Name,
		PlanSettings: planResponse,
	}, nil
}

func (s *planService) Update(ctx context.Context, planId string, req dto.PlanUpdateRequest) error {
	plan, err := s.planRepo.FindByID(ctx, nil, planId)
	if err != nil {
		return err
	}

	plan.Name = req.Name

	err = s.planRepo.Update(ctx, nil, plan)
	if err != nil {
		return err
	}

	return nil
}

func (s *planService) Delete(ctx context.Context, planId string) error {
	plan, err := s.planRepo.FindByID(ctx, nil, planId)
	if err != nil {
		return err
	}

	err = s.planRepo.Delete(ctx, nil, plan.ID.String())
	if err != nil {
		return err
	}

	return nil
}
