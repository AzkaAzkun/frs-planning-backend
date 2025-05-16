package repository

import (
	"context"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"

	"gorm.io/gorm"
)

type (
	PlanRepository interface {
		Create(ctx context.Context, tx *gorm.DB, plan entity.Plan) (entity.Plan, error)
		FindAll(ctx context.Context, tx *gorm.DB, workspaceId string, metareq meta.Meta) (dto.GetAllPlan, error)
		FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Plan, error)
		Update(ctx context.Context, tx *gorm.DB, plan entity.Plan) error
		Delete(ctx context.Context, tx *gorm.DB, id string) error
	}

	planRepository struct {
		db *gorm.DB
	}
)

func NewPlanRepository(db *gorm.DB) PlanRepository {
	return &planRepository{db}
}

func (r *planRepository) Create(ctx context.Context, tx *gorm.DB, plan entity.Plan) (entity.Plan, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&plan).Error; err != nil {
		return entity.Plan{}, err
	}
	return plan, nil
}

func (r *planRepository) FindAll(ctx context.Context, tx *gorm.DB, workspaceId string, metareq meta.Meta) (dto.GetAllPlan, error) {
	if tx == nil {
		tx = r.db
	}

	var plans []entity.Plan

	tx = tx.WithContext(ctx).Model(&entity.Plan{}).Where("workspace_id = ?", workspaceId)
	if err := WithFilters(tx, &metareq, AddModels(entity.Plan{})).Find(&plans).Error; err != nil {
		return dto.GetAllPlan{}, err
	}

	return dto.GetAllPlan{
		Plans: plans,
		Meta:  metareq,
	}, nil
}

func (r *planRepository) FindByID(ctx context.Context, tx *gorm.DB, id string) (entity.Plan, error) {
	if tx == nil {
		tx = r.db
	}

	var plan entity.Plan

	if err := tx.WithContext(ctx).Where("id = ?", id).Preload("PlanSettings").Preload("PlanSettings.Class").First(&plan).Error; err != nil {
		return entity.Plan{}, err
	}

	return plan, nil
}

func (r *planRepository) Update(ctx context.Context, tx *gorm.DB, plan entity.Plan) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&plan).Error; err != nil {
		return err
	}

	return nil
}

func (r *planRepository) Delete(ctx context.Context, tx *gorm.DB, id string) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("id = ?", id).Delete(&entity.Plan{}).Error; err != nil {
		return err
	}

	return nil
}
