package entity

import "github.com/google/uuid"

type ClassSettings struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Permission string    `json:"permission" gorm:"type:varchar(255)"`
	Used       int       `json:"used"`
}

func (ClassSettings) TableName() string {
	return "class_settings"
}
