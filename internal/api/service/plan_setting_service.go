package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	myerror "frs-planning-backend/internal/pkg/error"
	"net/http"

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
		classRepo       repository.ClassRepository
		planService     PlanService
		db              *gorm.DB
	}
)

func NewPlanSettingService(planSettingRepo repository.PlanSettingRepository,
	classRepo repository.ClassRepository,
	planService PlanService,
	db *gorm.DB) PlanSettingService {
	return &planSettingService{
		planSettingRepo: planSettingRepo,
		classRepo:       classRepo,
		planService:     planService,
		db:              db,
	}
}

func (s *planSettingService) Create(ctx context.Context, req dto.PlanSettingCreateRequest) error {
	plan, err := s.planService.GetPlanDetail(ctx, req.PlanID)
	if err != nil {
		return err
	}

	newclass, err := s.classRepo.FindByID(ctx, nil, req.ClassID)
	if err != nil {
		return err
	}

	for _, planSetting := range plan.PlanSettings {
		class := planSetting.Class
		if class.ID == req.ClassID {
			return myerror.New("class already exists in this plan", http.StatusConflict)
		}

		if class.CourseID == newclass.CourseID.String() {
			return myerror.New("class already exists in this course", http.StatusConflict)
		}

		if newclass.Day == class.Day {
			if newclass.StartTime >= class.StartTime && newclass.StartTime <= class.EndTime {
				return myerror.New("class time conflict", http.StatusConflict)
			}
			if newclass.EndTime >= class.StartTime && newclass.EndTime <= class.EndTime {
				return myerror.New("class time conflict", http.StatusConflict)
			}
		}
	}

	_, err = s.planSettingRepo.Create(ctx, nil, entity.PlanSettings{
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
