package service

import (
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"time"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	ClassSettingService interface {
		Create(ctx context.Context, req dto.CreateClassSettingRequest, userid string) (dto.ClassSettingResponse, error)
		Clone(ctx context.Context, userid string, req dto.CloneClassSettingRequest) (dto.ClassSettingResponse, error)
	}

	classSettingService struct {
		classSettingRepository repository.ClassSettingRepository
		db                     *gorm.DB
	}
)

func NewClassSettingService(classSettingRepository repository.ClassSettingRepository, db *gorm.DB) ClassSettingService {
	return &classSettingService{
		classSettingRepository: classSettingRepository,
		db:                     db,
	}
}

func (s *classSettingService) Create(ctx context.Context, req dto.CreateClassSettingRequest, userid string) (dto.ClassSettingResponse, error) {

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	classSetting, err := s.classSettingRepository.Create(ctx, nil, entity.ClassSettings{
		Name:       req.Name,
		Permission: req.Permission,
		UserID:     uuid.MustParse(userid),
		Status:     "CLONE",
	})
	if err != nil {
		return dto.ClassSettingResponse{}, nil
	}

	if err := tx.Commit().Error; err != nil {
		return dto.ClassSettingResponse{}, err
	}

	return dto.ClassSettingResponse{
		ID: classSetting.ID.String(),
		// Name field removed as it no longer exists in the entity
		Name:       classSetting.Name,
		User_id:    classSetting.UserID.String(),
		Permission: classSetting.Permission,
		Status:     classSetting.Status,
	}, nil
}

func parseTime(timeStr string) time.Time {
	t, _ := time.Parse(time.RFC3339, timeStr)
	return t
}

func (s *classSettingService) Clone(ctx context.Context, userid string, req dto.CloneClassSettingRequest) (dto.ClassSettingResponse, error) {
	cloneClassSetting, err := s.classSettingRepository.Clone(ctx, nil, uuid.MustParse(userid), uuid.MustParse(req.ClassSettingId))
	if err != nil {
		return dto.ClassSettingResponse{}, err
	}
	return dto.ClassSettingResponse{
		ID: cloneClassSetting.ID.String(),

		// Name field removed as it no longer exists in the entity
		User_id:    cloneClassSetting.UserID.String(),
		Permission: cloneClassSetting.Permission,

		Name:   cloneClassSetting.Name,
		Status: cloneClassSetting.Status,
	}, nil
}
