package repository

import (
	"gorm.io/gorm"
)

type (
	PlanRepository interface {
		// Create(ctx context.Context, tx *gorm.DB, plan entity.Plan) (entity.Plan, error)
		// FindAll(ctx context.Context, tx *gorm.DB, metareq meta.Meta) ([]entity.Plan, error)
		// FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Plan, error)
		// Update(ctx context.Context, tx *gorm.DB, plan entity.Plan) error
		// Delete(ctx context.Context, tx *gorm.DB, id string) error
	}

	planRepository struct {
		db *gorm.DB
	}
)

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db}
}
