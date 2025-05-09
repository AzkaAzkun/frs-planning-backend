package repository

import (
	"errors"
	"frs-planning-backend/internal/entity"
	"gorm.io/gorm"
)

type ClassRepository interface {
	Create(class *entity.Class) error
	FindAll() ([]entity.Class, error)
	FindByID(id int64) (*entity.Class, error)
	Update(class *entity.Class) error
	Delete(id int64) error
}

type classRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) ClassRepository {
	return &classRepository{db: db}
}

func (r *classRepository) Create(class *entity.Class) error {
	if err := r.db.Create(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) FindAll() ([]entity.Class, error) {
	var classes []entity.Class
	if err := r.db.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

func (r *classRepository) FindByID(id int64) (*entity.Class, error) {
	var class entity.Class
	if err := r.db.First(&class, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &class, nil
}

func (r *classRepository) Update(class *entity.Class) error {
	if err := r.db.Save(class).Error; err != nil {
		return err
	}
	return nil
}

func (r *classRepository) Delete(id int64) error {
	if err := r.db.Delete(&entity.Class{}, id).Error; err != nil {
		return err
	}
	return nil
}
