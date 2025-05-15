package service

import (
	"context"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	WorkspaceService interface {
		Create(ctx context.Context, req dto.CreateWorkspaceRequest, userid string) (dto.WorkspaceResponse, error)
		Find(ctx context.Context, workspaceid uuid.UUID) (dto.WorkspaceResponse, error)
		Get(ctx context.Context, userid string, metaReq meta.Meta) (dto.GetAllWorkspaceResponse, error)
		Update(ctx context.Context, req dto.UpdateWorkspaceRequest) (dto.WorkspaceResponse, error)
		Delete(ctx context.Context, req string) (dto.WorkspaceResponse, error)
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

func (s *workspaceService) Create(ctx context.Context, req dto.CreateWorkspaceRequest, userid string) (dto.WorkspaceResponse, error) {

	workspace, err := s.workspaceRepo.Create(ctx, nil, entity.Workspace{
		Name:           req.Name,
		ClassSettingID: nil,
	}, uuid.MustParse(userid))
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

func (s *workspaceService) Get(ctx context.Context, userid string, metaReq meta.Meta) (dto.GetAllWorkspaceResponse, error) {
	userWorkspaces, err := s.workspaceRepo.Get(ctx, nil, uuid.MustParse(userid), metaReq)
	if err != nil {
		return dto.GetAllWorkspaceResponse{}, err
	}

	// Map only id and name to the DTO
	var responses []dto.WorkspaceResponse
	for _, w := range userWorkspaces.Workspaces {
		responses = append(responses, dto.WorkspaceResponse{
			ID:   w.ID.String(),
			Name: w.Name,
		})
	}

	return dto.GetAllWorkspaceResponse{
		Workspaces: responses,
		Meta:       userWorkspaces.Meta,
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

func (s *workspaceService) Delete(ctx context.Context, id string) (dto.WorkspaceResponse, error) {
	workspaceUUID, err := uuid.Parse(id)
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
