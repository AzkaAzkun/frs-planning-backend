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
	FilmListService interface {
		Create(ctx context.Context, req dto.FilmListRequest, userId string) error
		UpdateVisibility(ctx context.Context, req dto.FilmListVisibilityRequest, filmlistId string) error
	}

	filmListService struct {
		filmListRepository repository.FilmListRepository
		filmRepository     repository.FilmRepository
		reviewRepository   repository.ReviewRepository
		db                 *gorm.DB
	}
)

func NewFilmList(filmListRepository repository.FilmListRepository,
	filmRepository repository.FilmRepository,
	reviewRepository repository.ReviewRepository,
	db *gorm.DB) FilmListService {
	return &filmListService{
		filmListRepository: filmListRepository,
		filmRepository:     filmRepository,
		reviewRepository:   reviewRepository,
		db:                 db,
	}
}

func (s *filmListService) Create(ctx context.Context, req dto.FilmListRequest, userId string) error {
	film, err := s.filmRepository.GetById(ctx, nil, req.FilmId)
	if err != nil {
		return err
	}

	review, err := s.reviewRepository.GetByFilmId(ctx, nil, film.ID.String())
	if err != nil {
		return err
	}

	if len(review) > 0 && req.ListStatus == string(entity.ListStatusPlanToWatch) {
		return myerror.New("film has been review", http.StatusBadRequest)
	}

	if req.ListStatus != string(entity.ListStatusPlanToWatch) && film.AiringStatus == entity.NotYetAired {
		return myerror.New("film is not aired yet", http.StatusBadRequest)
	}

	_, err = s.filmListRepository.Create(ctx, nil, entity.FilmList{
		FilmId:     film.ID,
		UserId:     uuid.MustParse(userId),
		ListStatus: entity.ListStatus(req.ListStatus),
		Visibility: entity.VisibilityPublic,
	})

	return err
}

func (s *filmListService) UpdateVisibility(ctx context.Context, req dto.FilmListVisibilityRequest, filmlistId string) error {
	filmlist, err := s.filmListRepository.GetById(ctx, nil, filmlistId)
	if err != nil {
		return err
	}

	filmlist.Visibility = entity.Visibility(req.Visibility)

	_, err = s.filmListRepository.Update(ctx, nil, filmlist)
	if err != nil {
		return err
	}

	return nil
}
