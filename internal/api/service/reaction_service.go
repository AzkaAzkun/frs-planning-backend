package service

import (
	"context"
	"film-management-api-golang/internal/api/repository"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	ReactionService interface {
		Create(ctx context.Context, req dto.ReactionRequest, userId string) error
		Update(ctx context.Context, req dto.ReactionUpdate, reactionId string) error
	}

	reactionService struct {
		reactionRepository repository.ReactionRepository
		reviewRepository   repository.ReviewRepository
		db                 *gorm.DB
	}
)

func NewReaction(reactionRepository repository.ReactionRepository,
	reviewRepository repository.ReviewRepository,
	db *gorm.DB) ReactionService {
	return &reactionService{
		reactionRepository: reactionRepository,
		reviewRepository:   reviewRepository,
		db:                 db,
	}
}

func (s *reactionService) Create(ctx context.Context, req dto.ReactionRequest, userId string) error {
	review, err := s.reviewRepository.GetById(ctx, nil, req.ReviewId)
	if err != nil {
		return err
	}

	_, err = s.reactionRepository.Create(ctx, nil, entity.Reaction{
		ReviewId: review.ID,
		UserId:   uuid.MustParse(userId),
		Status:   entity.ReactionStatus(req.Status),
	})
	if err != nil {
		return err
	}

	if req.Status == string(entity.StatusLike) {
		review.Likes++
	} else {
		review.Dislikes++
	}
	_, err = s.reviewRepository.Update(ctx, nil, review)
	if err != nil {
		return err
	}

	return nil
}

func (s *reactionService) Update(ctx context.Context, req dto.ReactionUpdate, reactionId string) error {
	reaction, err := s.reactionRepository.GetById(ctx, nil, reactionId)
	if err != nil {
		return err
	}

	if req.Status != string(reaction.Status) {
		review, err := s.reviewRepository.GetById(ctx, nil, reaction.ReviewId.String())
		if err != nil {
			return err
		}
		if req.Status == string(entity.StatusLike) {
			review.Likes++
			review.Dislikes--
		} else {
			review.Dislikes++
			review.Likes--
		}
		_, err = s.reviewRepository.Update(ctx, nil, review)
		if err != nil {
			return err
		}
	}

	reaction.Status = entity.ReactionStatus(req.Status)
	_, err = s.reactionRepository.Update(ctx, nil, reaction)
	if err != nil {
		return err
	}

	return err
}
