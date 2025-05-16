package repository

import (
	"context"
	"frs-planning-backend/internal/entity"

	"gorm.io/gorm"
)

type (
	PlanSettingRepository interface {
		Create(ctx context.Context, tx *gorm.DB, plan entity.PlanSettings) (entity.PlanSettings, error)
		Delete(ctx context.Context, tx *gorm.DB, id string) error
	}

	planSettingRepository struct {
		db *gorm.DB
	}
)

func NewPlanSettingRepository(db *gorm.DB) PlanSettingRepository {
	return &planSettingRepository{db}
}

func (r *planSettingRepository) Create(ctx context.Context, tx *gorm.DB, plan entity.PlanSettings) (entity.PlanSettings, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&plan).Error; err != nil {
		return entity.PlanSettings{}, err
	}
	return plan, nil
}

func (r *planSettingRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entity.PlanSettings{}).Error; err != nil {
		return err
	}
	return nil
}
