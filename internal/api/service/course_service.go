package service

import (
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
)

type CourseService interface {
	CreateCourse(req *dto.CreateCourseRequest) error
	GetAllCourses() ([]dto.CourseResponse, error)
	GetCourseByID(id string) (*dto.CourseResponse, error)
	UpdateCourse(id string, req *dto.UpdateCourseRequest) error
	DeleteCourse(id string) error
}

type courseService struct {
	courseRepo repository.CourseRepository
}

func NewCourseService(courseRepo repository.CourseRepository) CourseService {
	return &courseService{
		courseRepo: courseRepo,
	}
}

func (s *courseService) CreateCourse(req *dto.CreateCourseRequest) error {
	course := &entity.Course{
		Name:           req.Name,
		ClassSettingID: req.ClassSettingID,
	}

	return s.courseRepo.Create(course)
}

func (s *courseService) GetAllCourses() ([]dto.CourseResponse, error) {
	courses, err := s.courseRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.CourseResponse
	for _, course := range courses {
		responses = append(responses, dto.CourseResponse{
			ID:             course.ID,
			Name:           course.Name,
			ClassSettingID: course.ClassSettingID,
		})
	}

	return responses, nil
}

func (s *courseService) GetCourseByID(id string) (*dto.CourseResponse, error) {
	course, err := s.courseRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if course == nil {
		return nil, nil
	}

	return &dto.CourseResponse{
		ID:             course.ID,
		Name:           course.Name,
		ClassSettingID: course.ClassSettingID,
	}, nil
}

func (s *courseService) UpdateCourse(id string, req *dto.UpdateCourseRequest) error {
	course := &entity.Course{
		ID:             id,
		Name:           req.Name,
		ClassSettingID: req.ClassSettingID,
	}

	return s.courseRepo.Update(course)
}

func (s *courseService) DeleteCourse(id string) error {
	return s.courseRepo.Delete(id)
}
