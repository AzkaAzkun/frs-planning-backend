package service

import (
	"context"
	"fmt"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClassService interface {
	CreateClass(ctx context.Context, req dto.CreateClassRequest) (dto.ClassResponse, error)
	GetAllClasses(ctx context.Context, metareq meta.Meta) ([]dto.ClassResponse, error)
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
	courseID := uuid.MustParse(req.CourseID)
	course, err := s.courseRepo.FindByID(ctx, nil, req.CourseID)
	if err != nil {
		return dto.ClassResponse{}, err
	}

	day, err := parseWeekday(req.Day)
	if err != nil {
		return dto.ClassResponse{}, err
	}

	createResult, err := s.classRepo.Create(ctx, nil, entity.Class{
		Lecturer:  req.Lecturer,
		CourseID:  courseID,
		Name:      course.Name,
		Day:       day.String(),
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Classroom: req.Classroom,
	})
	if err != nil {
		return dto.ClassResponse{}, err
	}

	return dto.ClassResponse{
		ID:        createResult.ID.String(),
		Lecturer:  createResult.Lecturer,
		Day:       createResult.Day,
		Name:      createResult.Name,
		CourseID:  createResult.CourseID.String(),
		StartTime: createResult.StartTime,
		EndTime:   createResult.EndTime,
		Classroom: createResult.Classroom,
	}, nil
}

func (s *classService) GetAllClasses(ctx context.Context, metareq meta.Meta) ([]dto.ClassResponse, error) {
	classes, err := s.classRepo.FindAll(ctx, nil, metareq)
	if err != nil {
		return nil, err
	}

	var responses []dto.ClassResponse
	for _, class := range classes {
		responses = append(responses, dto.ClassResponse{
			ID:        class.ID.String(),
			Lecturer:  class.Lecturer,
			CourseID:  class.CourseID.String(),
			Name:      class.Name,
			Day:       class.Day,
			StartTime: class.StartTime,
			EndTime:   class.EndTime,
			Classroom: class.Classroom,
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
		ID:        class.ID.String(),
		Lecturer:  class.Lecturer,
		CourseID:  class.CourseID.String(),
		Day:       class.Day,
		Name:      class.Name,
		StartTime: class.StartTime,
		EndTime:   class.EndTime,
		Classroom: class.Classroom,
	}, nil
}

func (s *classService) UpdateClass(ctx context.Context, id string, req dto.UpdateClassRequest) error {
	course, err := s.courseRepo.FindByID(ctx, nil, req.CourseID)
	if err != nil {
		return err
	}

	day, err := parseWeekday(req.Day)
	if err != nil {
		return err
	}

	class := entity.Class{
		ID:        uuid.MustParse(id),
		Lecturer:  req.Lecturer,
		CourseID:  uuid.MustParse(req.CourseID),
		Day:       day.String(),
		Name:      course.Name,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Classroom: req.Classroom,
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
			ID:        class.ID.String(),
			Lecturer:  class.Lecturer,
			CourseID:  class.CourseID.String(),
			Day:       class.Day,
			StartTime: class.StartTime,
			EndTime:   class.EndTime,
			Classroom: class.Classroom,
		})
	}

	return responses, nil

}

func parseWeekday(day string) (time.Weekday, error) {
	switch strings.ToLower(day) {
	case "sunday":
		return time.Sunday, nil
	case "monday":
		return time.Monday, nil
	case "tuesday":
		return time.Tuesday, nil
	case "wednesday":
		return time.Wednesday, nil
	case "thursday":
		return time.Thursday, nil
	case "friday":
		return time.Friday, nil
	case "saturday":
		return time.Saturday, nil
	default:
		return time.Sunday, fmt.Errorf("invalid weekday: %s", day)
	}
}
