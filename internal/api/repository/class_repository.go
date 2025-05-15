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
	ClassRepository interface {
		Create(ctx context.Context, tx *gorm.DB, class entity.Class) (entity.Class, error)
		FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Class, error)
		FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Class, error)
		Update(ctx context.Context, tx *gorm.DB, class entity.Class) (entity.Class, error)
		Delete(ctx context.Context, tx *gorm.DB, id string) error

		FindByCourseID(ctx context.Context, tx *gorm.DB, courseID string) ([]entity.Class, error)

	}

	classRepository struct {
		db *gorm.DB
	}
)

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(ctx context.Context, tx *gorm.DB, class entity.Class) (entity.Class, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&class).Error; err != nil {
		return entity.Class{}, err
	}
	return class, nil
}

func (r *classRepository) FindAll(ctx context.Context, tx *gorm.DB) ([]entity.Class, error) {
	if tx == nil {
		tx = r.db
	}

	var classes []entity.Class
	if err := tx.WithContext(ctx).Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Class, error) {
	if tx == nil {
		tx = r.db
	}

	var class entity.Class

	// Use primary key query with where clause to avoid SQL parsing issues
	if err := tx.WithContext(ctx).Where("id = ?", id).First(&class).Error; err != nil {

	if err := tx.WithContext(ctx).First(&class, id).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.Class{}, myerror.New("class not found", http.StatusBadRequest)
		}
		return entity.Class{}, err
	}

	return class, nil
}

func (r *classRepository) Update(ctx context.Context, tx *gorm.DB, class entity.Class) (entity.Class, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&class).Error; err != nil {
		return entity.Class{}, err
	}
	return class, nil
}

func (r *classRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&entity.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) FindByCourseID(ctx context.Context, tx *gorm.DB, courseID string) ([]entity.Class, error) {
	if tx == nil {
		tx = r.db
	}

	var classes []entity.Class
	if err := tx.WithContext(ctx).Where("course_id = ?", courseID).Find(&classes).Error; err != nil {
		return nil, err
	}

	return classes, nil
}
