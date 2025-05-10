package entity

import "github.com/google/uuid"

type PlanSettings struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlanID  uuid.UUID `json:"plan_id" gorm:"type:uuid"`
	ClassID uuid.UUID `json:"class_id" gorm:"type:uuid"`
	Status  string    `json:"status" gorm:"type:varchar(255)"`
	IsLock  int64     `json:"is_lock"`
}

func (PlanSettings) TableName() string {
	return "plan_settings"
}
