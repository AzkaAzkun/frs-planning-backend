package repository

import (
	"errors"
	"frs-planning-backend/internal/entity"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *entity.Course) error
	FindAll() ([]entity.Course, error)
	FindByID(id string) (*entity.Course, error)
	Update(course *entity.Course) error
	Delete(id string) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(course *entity.Course) error {
	if err := r.db.Create(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course
	if err := r.db.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) FindByID(id string) (*entity.Course, error) {
	var course entity.Course
	if err := r.db.First(&course, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) Update(course *entity.Course) error {
	if err := r.db.Save(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) Delete(id string) error {
	if err := r.db.Delete(&entity.Course{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
