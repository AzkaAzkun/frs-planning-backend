package entity

import "github.com/google/uuid"

type Course struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name           string    `json:"name"`
	ClassSettingID uuid.UUID `json:"class_setting_id"`

	Timestamp
}

func (Course) TableName() string {
	return "courses"
}
