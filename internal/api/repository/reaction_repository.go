package repository

import (
	"context"
	"film-management-api-golang/internal/entity"

	"gorm.io/gorm"
)

type (
	ReactionRepository interface {
		Create(ctx context.Context, tx *gorm.DB, reaction entity.Reaction) (entity.Reaction, error)
		GetById(ctx context.Context, tx *gorm.DB, reactionId string) (entity.Reaction, error)
		Update(ctx context.Context, tx *gorm.DB, reaction entity.Reaction) (entity.Reaction, error)
	}

	reactionRepository struct {
		db *gorm.DB
	}
)

func NewReaction(db *gorm.DB) ReactionRepository {
	return &reactionRepository{db}
}

func (r *reactionRepository) Create(ctx context.Context, tx *gorm.DB, reaction entity.Reaction) (entity.Reaction, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&reaction).Error; err != nil {
		return reaction, err
	}

	return reaction, nil
}

func (r *reactionRepository) GetById(ctx context.Context, tx *gorm.DB, reactionId string) (entity.Reaction, error) {
	if tx == nil {
		tx = r.db
	}

	var reaction entity.Reaction
	if err := tx.WithContext(ctx).Where("id = ?", reactionId).Take(&reaction).Error; err != nil {
		return reaction, err
	}

	return reaction, nil
}

func (r *reactionRepository) Update(ctx context.Context, tx *gorm.DB, reaction entity.Reaction) (entity.Reaction, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&reaction).Error; err != nil {
		return reaction, err
	}

	return reaction, nil
}
