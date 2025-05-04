package repository

import (
	"context"
	"film-management-api-golang/internal/entity"

	"gorm.io/gorm"
)

type (
	FilmListRepository interface {
		Create(ctx context.Context, tx *gorm.DB, filmlist entity.FilmList) (entity.FilmList, error)
		GetById(ctx context.Context, tx *gorm.DB, filmlistId string) (entity.FilmList, error)
		Update(ctx context.Context, tx *gorm.DB, filmlist entity.FilmList) (entity.FilmList, error)
	}

	filmListRepository struct {
		db *gorm.DB
	}
)

func NewFilmList(db *gorm.DB) FilmListRepository {
	return &filmListRepository{
		db: db,
	}
}

func (r *filmListRepository) Create(ctx context.Context, tx *gorm.DB, filmlist entity.FilmList) (entity.FilmList, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&filmlist).Error; err != nil {
		return entity.FilmList{}, err
	}

	return filmlist, nil
}

func (r *filmListRepository) GetById(ctx context.Context, tx *gorm.DB, filmlistId string) (entity.FilmList, error) {
	if tx == nil {
		tx = r.db
	}

	var filmlist entity.FilmList
	if err := tx.WithContext(ctx).Where("id = ?", filmlistId).Take(&filmlist).Error; err != nil {
		return entity.FilmList{}, err
	}

	return filmlist, nil
}

func (r *filmListRepository) Update(ctx context.Context, tx *gorm.DB, filmlist entity.FilmList) (entity.FilmList, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&filmlist).Error; err != nil {
		return entity.FilmList{}, err
	}

	return filmlist, nil
}
