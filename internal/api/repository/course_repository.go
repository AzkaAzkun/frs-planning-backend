package repository

import (
	"context"
	"errors"
	"frs-planning-backend/internal/entity"
	myerror "frs-planning-backend/internal/pkg/error"
	"net/http"

	"gorm.io/gorm"
)

type (
	CourseRepository interface {
		Create(ctx context.Context, tx *gorm.DB, course entity.Course) (entity.Course, error)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Course, error)
		FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Course, error)
		Update(ctx context.Context, tx *gorm.DB, course entity.Course) error
		Delete(ctx context.Context, tx *gorm.DB, id string) error
	}

	courseRepository struct {
		db *gorm.DB
	}
)

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(ctx context.Context, tx *gorm.DB, course entity.Course) (entity.Course, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&course).Error; err != nil {
		return entity.Course{}, err
	}
	return course, nil
}

func (r *courseRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Course, error) {
	if tx == nil {
		tx = r.db
	}

	var courses []entity.Course
	if err := tx.WithContext(ctx).Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

func (r *courseRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Course, error) {
	if tx == nil {
		tx = r.db
	}

	var course entity.Course
	if err := tx.WithContext(ctx).First(&course, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Course{}, myerror.New("course not found", http.StatusBadRequest)
		}
		return entity.Course{}, err
	}
	return course, nil
}

func (r *courseRepository) Update(ctx context.Context, tx *gorm.DB, course entity.Course) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(course).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Course{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
