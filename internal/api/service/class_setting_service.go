package service

import (
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	ClassSettingService interface {
		Create(ctx context.Context, req dto.CreateClassSettingRequest, userid string) (dto.ClassSettingResponse, error)
		Clone(ctx context.Context, userid string, req dto.CloneClassSettingRequest) (dto.ClassSettingResponse, error)
		GetAll(ctx context.Context, metareq meta.Meta) (dto.ClassSettingListResponse, error)
		GetAllPrivate(ctx context.Context, userId string, metareq meta.Meta) (dto.ClassSettingListResponse, error)
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
		ID:         classSetting.ID.String(),
		Name:       classSetting.Name,
		User_id:    classSetting.UserID.String(),
		Permission: classSetting.Permission,
		Status:     classSetting.Status,
	}, nil
}

func (s *classSettingService) Clone(ctx context.Context, userid string, req dto.CloneClassSettingRequest) (dto.ClassSettingResponse, error) {
	cloneClassSetting, err := s.classSettingRepository.Clone(ctx, nil, uuid.MustParse(userid), uuid.MustParse(req.ClassSettingId))
	if err != nil {
		return dto.ClassSettingResponse{}, err
	}
	return dto.ClassSettingResponse{
		ID: cloneClassSetting.ID.String(),

		User_id:    cloneClassSetting.UserID.String(),
		Permission: cloneClassSetting.Permission,

		Name:   cloneClassSetting.Name,
		Status: cloneClassSetting.Status,
	}, nil
}

func (s *classSettingService) GetAll(ctx context.Context, metareq meta.Meta) (dto.ClassSettingListResponse, error) {
	classSettings, err := s.classSettingRepository.FindAll(ctx, nil, metareq)
	if err != nil {
		return dto.ClassSettingListResponse{}, err
	}

	var classSettingResponses []dto.ClassSettingResponse
	for _, classSetting := range classSettings.ClassSetting {
		classSettingResponses = append(classSettingResponses, dto.ClassSettingResponse{
			ID:         classSetting.ID.String(),
			Name:       classSetting.Name,
			User_id:    classSetting.UserID.String(),
			Permission: classSetting.Permission,
			Status:     classSetting.Status,
		})
	}

	return dto.ClassSettingListResponse{
		ClassSetting: classSettingResponses,
		Meta:         classSettings.Meta,
	}, nil
}

func (s *classSettingService) GetAllPrivate(ctx context.Context, userId string, metareq meta.Meta) (dto.ClassSettingListResponse, error) {
	classSettings, err := s.classSettingRepository.FindAllPrivate(ctx, nil, userId, metareq)
	if err != nil {
		return dto.ClassSettingListResponse{}, err
	}

	var classSettingResponses []dto.ClassSettingResponse
	for _, classSetting := range classSettings.ClassSetting {
		if classSetting.UserID.String() == userId {
			classSettingResponses = append(classSettingResponses, dto.ClassSettingResponse{
				ID:         classSetting.ID.String(),
				Name:       classSetting.Name,
				User_id:    classSetting.UserID.String(),
				Permission: classSetting.Permission,
				Status:     classSetting.Status,
			})
		}
	}

	return dto.ClassSettingListResponse{
		ClassSetting: classSettingResponses,
		Meta:         classSettings.Meta,
	}, nil
}
