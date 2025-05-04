package repository

import (
	"context"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	myerror "film-management-api-golang/internal/pkg/error"
	"film-management-api-golang/internal/pkg/meta"
	"net/http"

	"gorm.io/gorm"
)

type (
	GenreRepository interface {
		Create(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error)
		GetById(ctx context.Context, tx *gorm.DB, genreId string) (entity.Genre, error)
		GetByName(ctx context.Context, tx *gorm.DB, name string) (entity.Genre, error)
		GetBatchById(ctx context.Context, tx *gorm.DB, ids []string) ([]entity.Genre, error)
		GetAll(ctx context.Context, tx *gorm.DB) ([]dto.GenreResponse, error)
		GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GenreResponsePaginated, error)
		Update(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error)
	}

	genreRepository struct {
		db *gorm.DB
	}
)

func NewGenre(db *gorm.DB) GenreRepository {
	return &genreRepository{db}
}

func (r *genreRepository) Create(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&genre).Error; err != nil {
		return entity.Genre{}, err
	}

	return genre, nil
}

func (r *genreRepository) GetById(ctx context.Context, tx *gorm.DB, genreId string) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	var genre entity.Genre
	if err := tx.WithContext(ctx).Where("id = ?", genreId).Take(&genre).Error; err != nil {
		return entity.Genre{}, err
	}

	return genre, nil
}

func (r *genreRepository) GetByName(ctx context.Context, tx *gorm.DB, name string) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	var genre entity.Genre
	if err := tx.WithContext(ctx).Where("name = ?", name).Take(&genre).Error; err != nil {
		return entity.Genre{}, err
	}

	return genre, nil
}

func (r *genreRepository) GetBatchById(ctx context.Context, tx *gorm.DB, ids []string) ([]entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	var genres []entity.Genre
	if err := tx.WithContext(ctx).Where("id IN ?", ids).Find(&genres).Error; err != nil {
		return genres, err
	}

	if len(ids) != len(genres) {
		return genres, myerror.New("genres return not same with id request", http.StatusBadRequest)
	}

	return genres, nil
}

func (r *genreRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]dto.GenreResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var genres []dto.GenreResponse
	if err := tx.WithContext(ctx).Model(entity.Genre{}).Select("id", "name").Scan(&genres).Error; err != nil {
		return genres, err
	}

	return genres, nil
}

func (r *genreRepository) GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GenreResponsePaginated, error) {
	if tx == nil {
		tx = r.db
	}

	var genres []dto.GenreResponse
	tx = tx.WithContext(ctx).Model(entity.Genre{}).Select("id", "name")
	if err := WithFilters(tx, &metareq, AddModels(entity.Genre{})).Scan(&genres).Error; err != nil {
		return dto.GenreResponsePaginated{}, err
	}

	return dto.GenreResponsePaginated{
		Data: genres,
		Meta: metareq,
	}, nil
}

func (r *genreRepository) Update(ctx context.Context, tx *gorm.DB, genre entity.Genre) (entity.Genre, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&genre).Error; err != nil {
		return entity.Genre{}, err
	}

	return genre, nil
}
