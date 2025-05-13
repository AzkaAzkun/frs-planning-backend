package entity

import "github.com/google/uuid"

type PermissionClassSetting string

const (
	ClassSettingPublic  PermissionClassSetting = "PUBLIC"
	ClassSettingPrivate PermissionClassSetting = "PRIVATE"
)

type ClassSettings struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id"`
	Permission string    `json:"permission"`
	Used       int       `json:"used"`

	Timestamp
}

func (ClassSettings) TableName() string {
	return "class_settings"
}
