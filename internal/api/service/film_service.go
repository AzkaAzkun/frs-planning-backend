package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/pkg/meta"
	"film-management-api-golang/internal/utils"
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type (
	FilmService interface {
		Create(ctx context.Context, req dto.FilmCreateRequest) (dto.FilmCreateResponse, error)
		GetListFilm(ctx context.Context, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error)
		GetDetailFilm(ctx context.Context, filmId string) (dto.GetDetailFilmResponse, error)
	}

	filmService struct {
		filmRepository  repository.FilmRepository
		genreRepository repository.GenreRepository
		db              *gorm.DB
	}
)

func NewFilm(filmRepository repository.FilmRepository,
	genreRepository repository.GenreRepository,
	db *gorm.DB) FilmService {
	return &filmService{
		filmRepository:  filmRepository,
		genreRepository: genreRepository,
		db:              db,
	}
}

func (s *filmService) Create(ctx context.Context, req dto.FilmCreateRequest) (dto.FilmCreateResponse, error) {
	genreId := strings.Split(req.Genres, ",")

	genres, err := s.genreRepository.GetBatchById(ctx, nil, genreId)
	if err != nil {
		return dto.FilmCreateResponse{}, err
	}

	var creategenre []entity.FilmGenre
	for _, genre := range genres {
		creategenre = append(creategenre, entity.FilmGenre{
			GenreId: genre.ID,
		})
	}

	var createimage []entity.FilmImage
	for _, image := range req.Images {
		filename := fmt.Sprintf("film-%s-%s.%s", utils.ToSlug(req.Title), ulid.Make(), utils.GetExtensions(image.Filename))
		if err := utils.UploadFile(image, filename); err != nil {
			return dto.FilmCreateResponse{}, err
		}
		createimage = append(createimage, entity.FilmImage{
			ImagePath: filename,
		})
	}

	createResult, err := s.filmRepository.Create(ctx, nil, entity.Film{
		Title:         req.Title,
		Synopsis:      req.Synopsis,
		AiringStatus:  entity.AiringStatus(req.AiringStatus),
		TotalEpisodes: req.TotalEpisodes,
		ReleaseDate:   req.ReleaseDate,
		Images:        createimage,
		Genres:        creategenre,
	})
	if err != nil {
		return dto.FilmCreateResponse{}, err
	}

	return dto.FilmCreateResponse{
		ID: createResult.ID.String(),
	}, nil
}

func (s *filmService) GetListFilm(ctx context.Context, metareq meta.Meta) (dto.GetAllFilmPaginatedResponse, error) {
	data, err := s.filmRepository.GetAllPaginated(ctx, nil, metareq)
	if err != nil {
		return dto.GetAllFilmPaginatedResponse{}, err
	}

	return data, err
}

func (s *filmService) GetDetailFilm(ctx context.Context, filmId string) (dto.GetDetailFilmResponse, error) {
	data, err := s.filmRepository.GetDetailFilm(ctx, nil, filmId)
	if err != nil {
		return dto.GetDetailFilmResponse{}, err
	}

	var imagespath []string
	for _, image := range data.Images {
		imagespath = append(imagespath, image.ImagePath)
	}

	var genreresponse []dto.GenreResponse
	for _, genre := range data.Genres {
		genreresponse = append(genreresponse, dto.GenreResponse{
			ID:   genre.Genre.ID.String(),
			Name: genre.Genre.Name,
		})
	}

	return dto.GetDetailFilmResponse{
		ID:            data.ID,
		Title:         data.Title,
		Synopsis:      data.Synopsis,
		AiringStatus:  data.AiringStatus,
		TotalEpisodes: data.TotalEpisodes,
		ReleaseDate:   data.ReleaseDate.Format("2006-01-02 15:04:05"),
		Images:        imagespath,
		Genres:        genreresponse,
		AverageRating: data.AverageRating,
	}, nil
}
