package repository

import (
	"context"
	"film-management-api-golang/internal/entity"

	"gorm.io/gorm"
)

type (
	ReviewRepository interface {
		Create(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error)
		GetById(ctx context.Context, tx *gorm.DB, reviewId string) (entity.Review, error)
		GetByFilmId(ctx context.Context, tx *gorm.DB, filmId string) ([]entity.Review, error)
		Update(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error)
	}

	reviewRepository struct {
		db *gorm.DB
	}
)

func NewReview(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) Create(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&review).Error; err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) GetById(ctx context.Context, tx *gorm.DB, reviewId string) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	var review entity.Review
	if err := tx.WithContext(ctx).Where("id = ?", reviewId).Find(&review).Error; err != nil {
		return review, err
	}

	return review, nil
}

func (r *reviewRepository) GetByFilmId(ctx context.Context, tx *gorm.DB, filmId string) ([]entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	var reviews []entity.Review
	if err := tx.WithContext(ctx).Where("film_id = ?", filmId).Find(&reviews).Error; err != nil {
		return reviews, err
	}

	return reviews, nil
}

func (r *reviewRepository) Update(ctx context.Context, tx *gorm.DB, review entity.Review) (entity.Review, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&review).Error; err != nil {
		return review, err
	}

	return review, nil
}
