package service

import (
	"fmt"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	WorskspaceCollaboratorService interface {
		Add(ctx context.Context, req dto.AddCollaboratorRequest) (dto.CollaboratorResponse, error)
		Get(ctx context.Context, workspaceid uuid.UUID) ([]dto.CollaboratorResponse, error)
		Delete(ctx context.Context, req dto.DeleteCollaboratorRequest) (dto.CollaboratorResponse, error)
	}

	workspaceCollaboratorService struct {
		userRepository                  repository.UserRepository
		workspaceCollaboratorRepository repository.WorkspaceCollaboratorRepository
		db                              *gorm.DB
	}
)

func NewWorkspaceCollaboratorService(workspaceCollaboratorRepository repository.WorkspaceCollaboratorRepository, userRepository repository.UserRepository, db *gorm.DB) WorskspaceCollaboratorService {
	return &workspaceCollaboratorService{
		workspaceCollaboratorRepository: workspaceCollaboratorRepository,
		userRepository:                  userRepository,
		db:                              db,
	}
}

func (s *workspaceCollaboratorService) Add(ctx context.Context, req dto.AddCollaboratorRequest) (dto.CollaboratorResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.CollaboratorResponse{}, err
	}

	collab, err := s.workspaceCollaboratorRepository.Add(ctx, nil, entity.WorkspaceCollaborator{
		UserID:      user.ID,
		WorkspaceID: uuid.MustParse(req.Workspaceid),
		IsVerified:  req.IsVerified,
		Permission:  entity.CollabolatorPermission(req.Permission),
	})
	if err != nil {
		return dto.CollaboratorResponse{}, err
	}

	return dto.CollaboratorResponse{
		UserId:      collab.UserID.String(),
		Workspaceid: collab.WorkspaceID.String(),
		IsVerified:  collab.IsVerified,
		Permission:  string(collab.Permission),
	}, nil
}

func (s *workspaceCollaboratorService) Get(ctx context.Context, workspaceid uuid.UUID) ([]dto.CollaboratorResponse, error) {
	collabs, err := s.workspaceCollaboratorRepository.Get(ctx, nil, workspaceid)
	if err != nil {
		return nil, fmt.Errorf("failed to get collaborators: %w", err)
	}

	var responses []dto.CollaboratorResponse
	for _, c := range collabs {
		responses = append(responses, dto.CollaboratorResponse{
			UserId:      c.UserID.String(),
			Workspaceid: c.WorkspaceID.String(),
			IsVerified:  c.IsVerified,
			Permission:  string(c.Permission),
		})
	}

	return responses, nil
}

func (s *workspaceCollaboratorService) Delete(ctx context.Context, req dto.DeleteCollaboratorRequest) (dto.CollaboratorResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.CollaboratorResponse{}, err
	}

	workspaceUUID, err := uuid.Parse(req.Workspaceid)
	if err != nil {
		return dto.CollaboratorResponse{}, fmt.Errorf("invalid workspace ID: %w", err)
	}

	collab, err := s.workspaceCollaboratorRepository.Delete(ctx, nil, user.ID, workspaceUUID)
	if err != nil {
		return dto.CollaboratorResponse{}, err
	}

	return dto.CollaboratorResponse{
		UserId:      collab.UserID.String(),
		Workspaceid: collab.WorkspaceID.String(),
		IsVerified:  collab.IsVerified,
		Permission:  string(collab.Permission),
	}, nil
}
