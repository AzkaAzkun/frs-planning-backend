package entity

import "github.com/google/uuid"

type WorkspaceColaborator struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID      uuid.UUID `json:"user_id" gorm:"type:uuid"`
	WorkspaceID uuid.UUID `json:"workspace_id" gorm:"type:uuid"`
	Permission  string    `json:"permission" gorm:"type:varchar(255)"`
}

func (WorkspaceColaborator) TableName() string {
	return "workspace_colaborators"
}
