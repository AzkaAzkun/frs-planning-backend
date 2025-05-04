package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"
	myerror "film-management-api-golang/internal/pkg/error"
	"net/http"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ReviewService interface {
		Create(ctx context.Context, req dto.ReviewRequest, userId string) error
	}

	reviewService struct {
		reviewRepository repository.ReviewRepository
		filmRepository   repository.FilmRepository
		db               *gorm.DB
	}
)

func NewReview(reviewRepository repository.ReviewRepository,
	filmRepository repository.FilmRepository,
	db *gorm.DB) ReviewService {
	return &reviewService{
		reviewRepository: reviewRepository,
		filmRepository:   filmRepository,
		db:               db,
	}
}

func (s *reviewService) Create(ctx context.Context, req dto.ReviewRequest, userId string) error {
	film, err := s.filmRepository.GetById(ctx, nil, req.FilmId)
	if err != nil {
		return err
	}

	if film.AiringStatus == entity.NotYetAired {
		return myerror.New("not yet aired status film not allowed to review", http.StatusBadRequest)
	}

	_, err = s.reviewRepository.Create(ctx, nil, entity.Review{
		FilmId:  film.ID,
		UserId:  uuid.MustParse(userId),
		Rating:  req.Rating,
		Comment: req.Comment,
	})
	if err != nil {
		return err
	}

	return nil
}
