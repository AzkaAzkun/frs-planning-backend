package entity

import "github.com/google/uuid"

type Workspace struct {
	ID             uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name           string     `json:"name"`
	ClassSettingID *uuid.UUID `json:"class_setting_id"`

	Timestamp
}

func (Workspace) TableName() string {
	return "workspaces"
}
