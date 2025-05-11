package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassService interface {
	CreateClass(ctx context.Context, req dto.CreateClassRequest) (dto.ClassResponse, error)
	GetAllClasses(ctx context.Context) ([]dto.ClassResponse, error)
	GetClassByID(ctx context.Context, id string) (dto.ClassResponse, error)
	UpdateClass(ctx context.Context, id string, req dto.UpdateClassRequest) error
	DeleteClass(ctx context.Context, id string) error
}

type classService struct {
	classRepo repository.ClassRepository
	db        *gorm.DB
}

func NewClassService(classRepo repository.ClassRepository, db *gorm.DB) ClassService {
	return &classService{
		classRepo: classRepo,
		db:        db,
	}
}

func (s *classService) CreateClass(ctx context.Context, req dto.CreateClassRequest) (dto.ClassResponse, error) {
	createResult, err := s.classRepo.Create(ctx, nil, entity.Class{
		Lecturer:      req.Lecturer,
		CourseID:      uuid.MustParse(req.CourseID),
		ClassSchedule: req.ClassSchedule,
		Priority:      req.Priority,
	})
	if err != nil {
		return dto.ClassResponse{}, err
	}

	return dto.ClassResponse{
		ID:       createResult.ID.String(),
		Lecturer: createResult.Lecturer,
		CourseID: createResult.CourseID.String(),
		Priority: createResult.Priority,
	}, nil
}

func (s *classService) GetAllClasses(ctx context.Context) ([]dto.ClassResponse, error) {
	classes, err := s.classRepo.FindAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	var responses []dto.ClassResponse
	for _, class := range classes {
		responses = append(responses, dto.ClassResponse{
			ID:            class.ID.String(),
			Lecturer:      class.Lecturer,
			CourseID:      class.CourseID.String(),
			ClassSchedule: class.ClassSchedule,
			Priority:      class.Priority,
		})
	}

	return responses, nil
}

func (s *classService) GetClassByID(ctx context.Context, id string) (dto.ClassResponse, error) {
	class, err := s.classRepo.FindByID(ctx, nil, id)
	if err != nil {
		return dto.ClassResponse{}, err
	}

	return dto.ClassResponse{
		ID:            class.ID.String(),
		Lecturer:      class.Lecturer,
		CourseID:      class.CourseID.String(),
		ClassSchedule: class.ClassSchedule,
		Priority:      class.Priority,
	}, nil
}

func (s *classService) UpdateClass(ctx context.Context, id string, req dto.UpdateClassRequest) error {
	class := entity.Class{
		ID:            uuid.MustParse(id),
		Lecturer:      req.Lecturer,
		CourseID:      uuid.MustParse(req.CourseID),
		ClassSchedule: req.ClassSchedule,
		Priority:      req.Priority,
	}

	_, err := s.classRepo.Update(ctx, nil, class)
	if err != nil {
		return err
	}

	return nil
}

func (s *classService) DeleteClass(ctx context.Context, id string) error {
	return s.classRepo.Delete(ctx, nil, id)
}
