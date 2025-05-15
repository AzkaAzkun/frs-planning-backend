package repository

import (
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/entity"
	"frs-planning-backend/internal/pkg/meta"

	"github.com/google/uuid"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type (
	WorkspaceRepository interface {
		Create(ctx context.Context, tx *gorm.DB, workspace entity.Workspace, userid uuid.UUID) (entity.Workspace, error)
		Find(ctx context.Context, tx *gorm.DB, workspace uuid.UUID) (entity.Workspace, error)
		Get(ctx context.Context, tx *gorm.DB, userid uuid.UUID, metareq meta.Meta) (dto.GetAllWorkspace, error)
		Update(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID, name string) (entity.Workspace, error)
		Delete(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID) (entity.Workspace, error)
	}

	workspaceRepository struct {
		db *gorm.DB
	}
)

func NewWorkspaceRepository(db *gorm.DB) WorkspaceRepository {
	return &workspaceRepository{db}
}

func (r *workspaceRepository) Create(ctx context.Context, tx *gorm.DB, workspace entity.Workspace, userid uuid.UUID) (entity.Workspace, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&workspace).Error; err != nil {
		return entity.Workspace{}, err
	}

	// nambahin user yang buat workspace ke collaborator
	collab := entity.WorkspaceCollaborator{
		UserID:      userid,
		WorkspaceID: workspace.ID,
	}

	if err := tx.WithContext(ctx).Create(&collab).Error; err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) Find(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID) (entity.Workspace, error) {
	if tx == nil {
		tx = r.db
	}

	var workspace entity.Workspace

	if err := tx.WithContext(ctx).First(&workspace, "id = ?", workspaceid).Error; err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) Get(ctx context.Context, tx *gorm.DB, userid uuid.UUID, metareq meta.Meta) (dto.GetAllWorkspace, error) {
	if tx == nil {
		tx = r.db
	}
	var workspaceIDs []uuid.UUID
	if err := tx.WithContext(ctx).
		Model(&entity.WorkspaceCollaborator{}).
		Where("user_id = ?", userid).
		Pluck("workspace_id", &workspaceIDs).Error; err != nil {
		return dto.GetAllWorkspace{}, err
	}

	var userWorkspaces []entity.Workspace
	query := tx.WithContext(ctx).Model(entity.Workspace{}).Where("id IN ?", workspaceIDs)
	if err := WithFilters(query, &metareq, AddModels(entity.Workspace{})).Find(&userWorkspaces).Error; err != nil {
		return dto.GetAllWorkspace{}, err
	}

	return dto.GetAllWorkspace{
		Workspaces: userWorkspaces,
		Meta:       metareq,
	}, nil
}

func (r *workspaceRepository) Update(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID, name string) (entity.Workspace, error) {
	if tx == nil {
		tx = r.db
	}

	var workspace entity.Workspace

	if err := tx.WithContext(ctx).First(&workspace, "id = ?", workspaceid).Error; err != nil {
		return entity.Workspace{}, err
	}

	workspace.Name = name
	if err := tx.WithContext(ctx).Save(&workspace).Error; err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}

func (r *workspaceRepository) Delete(ctx context.Context, tx *gorm.DB, workspaceid uuid.UUID) (entity.Workspace, error) {
	if tx == nil {
		tx = r.db
	}

	var workspace entity.Workspace

	if err := tx.WithContext(ctx).First(&workspace, "id = ?", workspaceid).Error; err != nil {
		return entity.Workspace{}, err
	}

	// delete semua collaborator di workspace
	if err := tx.WithContext(ctx).Where("workspace_id = ?", workspaceid).Delete(&entity.WorkspaceCollaborator{}).Error; err != nil {
		return entity.Workspace{}, err
	}

	if err := tx.WithContext(ctx).Delete(&workspace).Error; err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}
