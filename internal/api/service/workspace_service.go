package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	WorkspaceService interface {
		Create(ctx context.Context, req dto.CreateWorkspaceRequest) (dto.WorkspaceResponse, error)
		Find(ctx context.Context, workspaceid uuid.UUID) (dto.WorkspaceResponse, error)
		Update(ctx context.Context, req dto.UpdateWorkspaceRequest) (dto.WorkspaceResponse, error)
		Delete(ctx context.Context, req dto.DeleteWorkspaceRequest) (dto.WorkspaceResponse, error)
	}

	workspaceService struct {
		workspaceRepo repository.WorkspaceRepository
		db            *gorm.DB
	}
)

func NewWorkspaceService(workRepository repository.WorkspaceRepository, db *gorm.DB) WorkspaceService {
	return &workspaceService{
		workspaceRepo: workRepository,
		db:            db,
	}
}

func (s *workspaceService) Create(ctx context.Context, req dto.CreateWorkspaceRequest) (dto.WorkspaceResponse, error) {
	userUUID, err := uuid.Parse(req.UserID)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}

	workspace, err := s.workspaceRepo.Create(ctx, nil, entity.Workspace{
		Name: req.Name,
		//class setting blom ada
	}, userUUID)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}

	return dto.WorkspaceResponse{
		Name: workspace.Name,
		ID:   workspace.ID.String(),
	}, nil
}

func (s *workspaceService) Find(ctx context.Context, workspaceid uuid.UUID) (dto.WorkspaceResponse, error) {

	workspace, err := s.workspaceRepo.Find(ctx, nil, workspaceid)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}

	return dto.WorkspaceResponse{
		ID:   workspace.ID.String(),
		Name: workspace.Name,
	}, nil
}

func (s *workspaceService) Update(ctx context.Context, req dto.UpdateWorkspaceRequest) (dto.WorkspaceResponse, error) {
	workspaceUUID, err := uuid.Parse(req.ID)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}
	workspace, err := s.workspaceRepo.Update(ctx, nil, workspaceUUID, req.Name)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}
	return dto.WorkspaceResponse{
		ID:   workspace.ID.String(),
		Name: workspace.Name,
	}, nil
}

func (s *workspaceService) Delete(ctx context.Context, req dto.DeleteWorkspaceRequest) (dto.WorkspaceResponse, error) {
	workspaceUUID, err := uuid.Parse(req.ID)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}
	workspace, err := s.workspaceRepo.Delete(ctx, nil, workspaceUUID)
	if err != nil {
		return dto.WorkspaceResponse{}, err
	}
	return dto.WorkspaceResponse{
		Name: workspace.Name,
		ID:   workspace.ID.String(),
	}, nil
}
