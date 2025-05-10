package service

import (
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
)

type ClassService interface {
	CreateClass(req *dto.CreateClassRequest) error
	GetAllClasses() ([]dto.ClassResponse, error)
	GetClassByID(id int64) (*dto.ClassResponse, error)
	UpdateClass(id int64, req *dto.UpdateClassRequest) error
	DeleteClass(id int64) error
}

type classService struct {
	classRepo repository.ClassRepository
}

func NewClassService(classRepo repository.ClassRepository) ClassService {
	return &classService{
		classRepo: classRepo,
	}
}

func (s *classService) CreateClass(req *dto.CreateClassRequest) error {
	class := &entity.Class{
		Lecturer:      req.Lecturer,
		CourseID:      req.CourseID,
		ClassSchedule: req.ClassSchedule,
		Priority:      req.Priority,
	}

	return s.classRepo.Create(class)
}

func (s *classService) GetAllClasses() ([]dto.ClassResponse, error) {
	classes, err := s.classRepo.FindAll()
	if err != nil {
		return nil, err
	}

	var responses []dto.ClassResponse
	for _, class := range classes {
		responses = append(responses, dto.ClassResponse{
			ID:            class.ID,
			Lecturer:      class.Lecturer,
			CourseID:      class.CourseID,
			ClassSchedule: class.ClassSchedule,
			Priority:      class.Priority,
		})
	}

	return responses, nil
}

func (s *classService) GetClassByID(id int64) (*dto.ClassResponse, error) {
	class, err := s.classRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if class == nil {
		return nil, nil
	}

	return &dto.ClassResponse{
		ID:            class.ID,
		Lecturer:      class.Lecturer,
		CourseID:      class.CourseID,
		ClassSchedule: class.ClassSchedule,
		Priority:      class.Priority,
	}, nil
}

func (s *classService) UpdateClass(id int64, req *dto.UpdateClassRequest) error {
	class := &entity.Class{
		ID:            id,
		Lecturer:      req.Lecturer,
		CourseID:      req.CourseID,
		ClassSchedule: req.ClassSchedule,
		Priority:      req.Priority,
	}

	return s.classRepo.Update(class)
}

func (s *classService) DeleteClass(id int64) error {
	return s.classRepo.Delete(id)
}
