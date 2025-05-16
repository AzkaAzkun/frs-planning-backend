package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	PlanSettingService interface {
		Create(ctx context.Context, req dto.PlanSettingCreateRequest) error
		Delete(ctx context.Context, planSettingId string) error
	}

	planSettingService struct {
		planSettingRepo repository.PlanSettingRepository
		db              *gorm.DB
	}
)

func NewPlanSettingService(planSettingRepo repository.PlanSettingRepository, db *gorm.DB) PlanSettingService {
	return &planSettingService{
		planSettingRepo: planSettingRepo,
		db:              db,
	}
}

func (s *planSettingService) Create(ctx context.Context, req dto.PlanSettingCreateRequest) error {
	_, err := s.planSettingRepo.Create(ctx, nil, entity.PlanSettings{
		PlanID:  uuid.MustParse(req.PlanID),
		ClassID: uuid.MustParse(req.ClassID),
	})
	if err != nil {
		return err
	}

	return err
}

func (s *planSettingService) Delete(ctx context.Context, planSettingId string) error {
	if err := s.planSettingRepo.Delete(ctx, nil, planSettingId); err != nil {
		return err
	}
	return nil
}
