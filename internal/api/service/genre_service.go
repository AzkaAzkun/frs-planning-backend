package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	myerror "film-management-api-golang/internal/pkg/error"
	"film-management-api-golang/internal/pkg/meta"
	"net/http"

	"gorm.io/gorm"
)

type (
	GenreService interface {
		Create(ctx context.Context, req dto.GenreRequest) (dto.GenreResponse, error)
		GetById(ctx context.Context, genreId string) (dto.GenreResponse, error)
		GetByName(ctx context.Context, name string) (dto.GenreResponse, error)
		GetAll(ctx context.Context) ([]dto.GenreResponse, error)
		GetAllPaginated(ctx context.Context, metareq meta.Meta) (dto.GenreResponsePaginated, error)
		Update(ctx context.Context, req dto.GenreRequest, genreId string) (dto.GenreResponse, error)
	}

	genreService struct {
		genreRepository repository.GenreRepository
		db              *gorm.DB
	}
)

func NewGenre(genreRepository repository.GenreRepository,
	db *gorm.DB) GenreService {
	return &genreService{
		genreRepository: genreRepository,
		db:              db,
	}
}

func (s *genreService) Create(ctx context.Context, req dto.GenreRequest) (dto.GenreResponse, error) {
	_, err := s.genreRepository.GetByName(ctx, nil, req.Name)
	if err == nil {
		return dto.GenreResponse{}, myerror.New("this genre already exist", http.StatusConflict)
	}

	createResult, err := s.genreRepository.Create(ctx, nil, entity.Genre{
		Name: req.Name,
	})
	if err != nil {
		return dto.GenreResponse{}, err
	}

	return dto.GenreResponse{
		ID:   createResult.ID.String(),
		Name: createResult.Name,
	}, nil
}

func (s *genreService) GetById(ctx context.Context, genreId string) (dto.GenreResponse, error) {
	genre, err := s.genreRepository.GetById(ctx, nil, genreId)
	if err != nil {
		return dto.GenreResponse{}, err
	}

	return dto.GenreResponse{
		ID:   genre.ID.String(),
		Name: genre.Name,
	}, nil
}

func (s *genreService) GetByName(ctx context.Context, name string) (dto.GenreResponse, error) {
	genre, err := s.genreRepository.GetByName(ctx, nil, name)
	if err != nil {
		return dto.GenreResponse{}, err
	}

	return dto.GenreResponse{
		ID:   genre.ID.String(),
		Name: genre.Name,
	}, nil
}

func (s *genreService) GetAll(ctx context.Context) ([]dto.GenreResponse, error) {
	genres, err := s.genreRepository.GetAll(ctx, nil)
	if err != nil {
		return []dto.GenreResponse{}, err
	}

	return genres, nil
}
func (s *genreService) GetAllPaginated(ctx context.Context, metareq meta.Meta) (dto.GenreResponsePaginated, error) {
	result, err := s.genreRepository.GetAllPaginated(ctx, nil, metareq)
	if err != nil {
		return dto.GenreResponsePaginated{}, err
	}

	return result, nil
}

func (s *genreService) Update(ctx context.Context, req dto.GenreRequest, genreId string) (dto.GenreResponse, error) {
	genre, err := s.genreRepository.GetById(ctx, nil, genreId)
	if err != nil {
		return dto.GenreResponse{}, err
	}

	same, err := s.genreRepository.GetByName(ctx, nil, req.Name)
	if err == nil && same.ID != genre.ID {
		return dto.GenreResponse{}, myerror.New("this genre already exist", http.StatusConflict)
	}

	genre.Name = req.Name

	updateResult, err := s.genreRepository.Update(ctx, nil, genre)
	if err != nil {
		return dto.GenreResponse{}, err
	}

	return dto.GenreResponse{
		ID:   updateResult.ID.String(),
		Name: updateResult.Name,
	}, nil
}
