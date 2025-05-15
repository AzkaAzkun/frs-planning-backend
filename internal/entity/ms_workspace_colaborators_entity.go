package entity

import "github.com/google/uuid"

type CollabolatorPermission string

const (
	CollabolatorReadOnly CollabolatorPermission = "READ_ONLY"
	CollabolatorEdit     CollabolatorPermission = "EDIT"
)

type WorkspaceCollaborator struct {
	ID          uuid.UUID              `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID              `json:"user_id" gorm:"type:uuid"`
	WorkspaceID uuid.UUID              `json:"workspace_id" gorm:"type:uuid"`
	IsVerified  bool                   `json:"is_verified"`
	Permission  CollabolatorPermission `json:"permission"`

	Timestamp
}

func (WorkspaceCollaborator) TableName() string {
	return "workspace_colaborators"
}
