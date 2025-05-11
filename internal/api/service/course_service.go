package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
)

type CourseService interface {
	CreateCourse(ctx context.Context, req dto.CreateCourseRequest) (dto.CourseResponse, error)
	GetAllCourses(ctx context.Context) ([]dto.CourseResponse, error)
	GetCourseByID(ctx context.Context, id string) (dto.CourseResponse, error)
	UpdateCourse(ctx context.Context, id string, req dto.UpdateCourseRequest) error
	DeleteCourse(ctx context.Context, id string) error
}

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
	}
}

func (s *courseService) CreateCourse(ctx context.Context, req dto.CreateCourseRequest) (dto.CourseResponse, error) {
	createResult, err := s.courseRepo.Create(ctx, nil, entity.Course{
		Name:           req.Name,
		ClassSettingID: uuid.MustParse(req.ClassSettingID),
	})
	if err != nil {
		return dto.CourseResponse{}, err
	}

	return dto.CourseResponse{
		ID:             createResult.ID.String(),
		Name:           createResult.Name,
		ClassSettingID: createResult.ClassSettingID.String(),
	}, nil
}

func (s *courseService) GetAllCourses(ctx context.Context) ([]dto.CourseResponse, error) {
	courses, err := s.courseRepo.FindAll(ctx, nil)
	if err != nil {
		return nil, err
	}

	var responses []dto.CourseResponse
	for _, course := range courses {
		responses = append(responses, dto.CourseResponse{
			ID:             course.ID.String(),
			Name:           course.Name,
			ClassSettingID: course.ClassSettingID.String(),
		})
	}

	return responses, nil
}

func (s *courseService) GetCourseByID(ctx context.Context, id string) (dto.CourseResponse, error) {
	course, err := s.courseRepo.FindByID(ctx, nil, id)
	if err != nil {
		return dto.CourseResponse{}, err
	}

	return dto.CourseResponse{
		ID:             course.ID.String(),
		Name:           course.Name,
		ClassSettingID: course.ClassSettingID.String(),
	}, nil
}

func (s *courseService) UpdateCourse(ctx context.Context, id string, req dto.UpdateCourseRequest) error {
	course := entity.Course{
		ID:             uuid.MustParse(id),
		Name:           req.Name,
		ClassSettingID: uuid.MustParse(req.ClassSettingID),
	}

	if err := s.courseRepo.Update(ctx, nil, course); err != nil {
		return err
	}

	return nil
}

func (s *courseService) DeleteCourse(ctx context.Context, id string) error {
	if err := s.courseRepo.Delete(ctx, nil, id); err != nil {
		return err
	}

	return nil
}
