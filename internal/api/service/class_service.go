package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassService interface {
	CreateClass(ctx context.Context, req dto.CreateClassRequest) (dto.ClassResponse, error)
	GetAllClasses(ctx context.Context) ([]dto.ClassResponse, error)
	GetClassByID(ctx context.Context, id string) (dto.ClassResponse, error)
	UpdateClass(ctx context.Context, id string, req dto.UpdateClassRequest) error
	DeleteClass(ctx context.Context, id string) error
	GetClassesByCourseID(ctx context.Context, courseID string) ([]dto.ClassResponse, error)
}

type classService struct {
	classRepo  repository.ClassRepository
	courseRepo repository.CourseRepository
	db         *gorm.DB
}

func NewClassService(classRepo repository.ClassRepository, courseRepo repository.CourseRepository, db *gorm.DB) ClassService {
	return &classService{
		classRepo:  classRepo,
		courseRepo: courseRepo,
		db:         db,
	}
}

func (s *classService) CreateClass(ctx context.Context, req dto.CreateClassRequest) (dto.ClassResponse, error) {
	// Check if the course exists before creating the class
	courseID := uuid.MustParse(req.CourseID)
	_, err := s.courseRepo.FindByID(ctx, nil, req.CourseID)
	if err != nil {
		return dto.ClassResponse{}, err
	}

	parsedTime, err := time.Parse("15:04:05", req.ClassSchedule)
	if err != nil {
		return dto.ClassResponse{}, err
	}

	createResult, err := s.classRepo.Create(ctx, nil, entity.Class{
		Lecturer:      req.Lecturer,
		CourseID:      courseID,
		ClassSchedule: parsedTime,
		Priority:      req.Priority,
		Classroom:     req.Classroom,
	})
	if err != nil {
		return dto.ClassResponse{}, err
	}

	return dto.ClassResponse{
		ID:        createResult.ID.String(),
		Lecturer:  createResult.Lecturer,
		CourseID:  createResult.CourseID.String(),
		Priority:  createResult.Priority,
		Classroom: createResult.Classroom,
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
			Classroom:     class.Classroom,
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
		Classroom:     class.Classroom,
	}, nil
}

func (s *classService) UpdateClass(ctx context.Context, id string, req dto.UpdateClassRequest) error {
	parsedTime, err := time.Parse("15:04:05", req.ClassSchedule)
	if err != nil {
		return err
	}

	class := entity.Class{
		ID:            uuid.MustParse(id),
		Lecturer:      req.Lecturer,
		CourseID:      uuid.MustParse(req.CourseID),
		ClassSchedule: parsedTime,
		Priority:      req.Priority,
	}

	_, err = s.classRepo.Update(ctx, nil, class)
	if err != nil {
		return err
	}

	return nil
}

func (s *classService) DeleteClass(ctx context.Context, id string) error {
	return s.classRepo.Delete(ctx, nil, id)
}

func (s *classService) GetClassesByCourseID(ctx context.Context, courseID string) ([]dto.ClassResponse, error) {
	classes, err := s.classRepo.FindByCourseID(ctx, nil, courseID)
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
			Classroom:     class.Classroom,
		})
	}

	return responses, nil
}
