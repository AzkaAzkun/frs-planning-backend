package repository

import (
	"frs-planning-backend/internal/entity"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	WorkspaceCollaboratorRepository interface {
		Add(ctx context.Context, tx *gorm.DB, collaborator entity.WorkspaceCollaborator) (entity.WorkspaceCollaborator, error)
		Get(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID) ([]CollaboratorWithUser, error)
		Delete(ctx context.Context, tx *gorm.DB, userid uuid.UUID, workspaceid uuid.UUID) (entity.WorkspaceCollaborator, error)
	}

	workspaceColllaboratorRepository struct {
		db *gorm.DB
	}

	CollaboratorWithUser struct {
		entity.User
		Permission string
		IsVerified bool
	}
)

func NewWOrkspaceCollaboratorRepository(db *gorm.DB) WorkspaceCollaboratorRepository {
	return &workspaceColllaboratorRepository{db}
}

func (r *workspaceColllaboratorRepository) Add(ctx context.Context, tx *gorm.DB, collaborator entity.WorkspaceCollaborator) (entity.WorkspaceCollaborator, error) {
	if tx == nil {
		tx = r.db
	}

	err := tx.WithContext(ctx).Create(&collaborator).Error
	if err != nil {
		return entity.WorkspaceCollaborator{}, err
	}

	return collaborator, nil
}

func (r *workspaceColllaboratorRepository) Get(ctx context.Context, tx *gorm.DB, workspaceID uuid.UUID) ([]CollaboratorWithUser, error) {
	if tx == nil {
		tx = r.db
	}

	var collaborators []entity.WorkspaceCollaborator
	err := tx.WithContext(ctx).
		Where("workspace_id = ?", workspaceID).
		Find(&collaborators).Error
	if err != nil {
		return nil, err
	}

	userIDs := make([]uuid.UUID, 0, len(collaborators))
	userMeta := make(map[uuid.UUID]struct {
		Permission string
		IsVerified bool
	})
	for _, c := range collaborators {
		userIDs = append(userIDs, c.UserID)
		userMeta[c.UserID] = struct {
			Permission string
			IsVerified bool
		}{
			Permission: string(c.Permission),
			IsVerified: c.IsVerified,
		}
	}

	var users []entity.User
	err = tx.WithContext(ctx).
		Where("id IN ?", userIDs).
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	// Combine data
	var result []CollaboratorWithUser
	for _, u := range users {
		meta := userMeta[u.ID]
		result = append(result, CollaboratorWithUser{
			User:       u,
			Permission: meta.Permission,
			IsVerified: meta.IsVerified,
		})
	}

	return result, nil
}

func (r *workspaceColllaboratorRepository) Delete(ctx context.Context, tx *gorm.DB, userid uuid.UUID, workspaceid uuid.UUID) (entity.WorkspaceCollaborator, error) {
	if tx == nil {
		tx = r.db
	}

	var collaborator entity.WorkspaceCollaborator
	err := tx.WithContext(ctx).
		Where("user_id = ? AND workspace_id = ?", userid, workspaceid).
		First(&collaborator).Error
	if err != nil {
		return entity.WorkspaceCollaborator{}, err
	}

	err = tx.WithContext(ctx).Delete(&collaborator).Error
	if err != nil {
		return entity.WorkspaceCollaborator{}, err
	}

	return collaborator, nil
}
