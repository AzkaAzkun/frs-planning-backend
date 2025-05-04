package repository

import (
	"context"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/pkg/meta"

	"gorm.io/gorm"
)

type (
	FilmRepository interface {
		Create(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error)
		GetById(ctx context.Context, tx *gorm.DB, filmId string) (entity.Film, error)
		GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error)
		GetDetailFilm(ctx context.Context, tx *gorm.DB, filmId string) (dto.GetDetailFilm, error)
	}

	filmRepository struct {
		db *gorm.DB
	}
)

func NewFilm(db *gorm.DB) FilmRepository {
	return &filmRepository{
		db: db,
	}
}

func (r *filmRepository) Create(ctx context.Context, tx *gorm.DB, film entity.Film) (entity.Film, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) GetById(ctx context.Context, tx *gorm.DB, filmId string) (entity.Film, error) {
	if tx == nil {
		tx = r.db
	}

	var film entity.Film
	if err := tx.WithContext(ctx).Where("id = ?", filmId).Take(&film).Error; err != nil {
		return entity.Film{}, err
	}

	return film, nil
}

func (r *filmRepository) GetAllPaginated(ctx context.Context, tx *gorm.DB, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error) {
	if tx == nil {
		tx = r.db
	}

	var result []dto.GetAllFilmResponse
	query := tx.WithContext(ctx).Model(&entity.Film{})
	query = WithFilters(query, &metareq, AddModels(entity.Film{}))
	subQuery := r.db.
		Select("film_id, ROUND(AVG(rating)::numeric, 2) as average_rating").
		Table("us_reviews").
		Group("film_id")
	query = query.
		Select("films.*, avg_ratings.average_rating").
		Joins("LEFT JOIN (?) as avg_ratings ON avg_ratings.film_id::uuid = films.id", subQuery).Scan(&result)

	return dto.GetAllFilmPaginatedResponse{
		Data: result,
		Meta: metareq,
	}, query.Error
}

func (r *filmRepository) GetDetailFilm(ctx context.Context, tx *gorm.DB, filmId string) (dto.GetDetailFilm, error) {
	if tx == nil {
		tx = r.db
	}

	var film dto.FilmWithRating

	subQuery := r.db.
		Select("film_id, ROUND(AVG(rating)::numeric, 2) as average_rating").
		Table("us_reviews").
		Group("film_id")

	err := tx.WithContext(ctx).
		Model(&entity.Film{}).
		Select("films.*, avg_ratings.average_rating").
		Joins("LEFT JOIN (?) as avg_ratings ON avg_ratings.film_id::uuid = films.id", subQuery).
		Where("films.id = ?", filmId).
		Preload("Images").
		Preload("Genres.Genre").
		First(&film).Error

	if err != nil {
		return dto.GetDetailFilm{}, err
	}

	result := dto.GetDetailFilm{
		ID:            film.ID.String(),
		Title:         film.Title,
		Synopsis:      film.Synopsis,
		AiringStatus:  string(film.AiringStatus),
		TotalEpisodes: film.TotalEpisodes,
		ReleaseDate:   film.ReleaseDate,
		Images:        film.Images,
		Genres:        film.Genres,
		AverageRating: film.AverageRating,
	}

	return result, nil
}
