package entity

import "github.com/google/uuid"

type PermissionClassSetting string
type PermissionClassStatus string

const (
	ClassSettingPublic  PermissionClassSetting = "PUBLIC"
	ClassSettingPrivate PermissionClassSetting = "PRIVATE"
)

const (
	ClassSettingOwn   PermissionClassStatus = "OWN"
	ClassSettingClone PermissionClassStatus = "CLONE"
)

type ClassSettings struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id"`
	Permission string    `json:"permission"`
	Used       int       `json:"used"`
	Status     string    `json:"status"`
	Name       string    `json:"name"`

	Timestamp
}

func (ClassSettings) TableName() string {
	return "class_settings"
}
