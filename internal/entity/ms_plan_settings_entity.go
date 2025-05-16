package entity

import "github.com/google/uuid"

type StatusPlanSettings string

const (
	PlanSettingSuccess StatusPlanSettings = "SUCCESS"
	PlanSettingFull    StatusPlanSettings = "FULL"
	PlanSettingPending StatusPlanSettings = "PENDING"
)

type PlanSettings struct {
	ID      uuid.UUID          `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	PlanID  uuid.UUID          `json:"plan_id" gorm:"type:uuid"`
	ClassID uuid.UUID          `json:"class_id" gorm:"type:uuid"`
	Status  StatusPlanSettings `json:"status"`
	IsLock  bool               `json:"is_lock"`

	Timestamp
}

func (PlanSettings) TableName() string {
	return "plan_settings"
}
