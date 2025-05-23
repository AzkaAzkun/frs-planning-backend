package entity

import "github.com/google/uuid"

type Plan struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	WorkspaceID uuid.UUID `json:"workspace_id" gorm:"type:uuid"`
	Name        string    `json:"name"`

	PlanSettings []PlanSettings `json:"plan_settings" gorm:"foreignKey:PlanID;references:ID"`

	Timestamp
}

func (Plan) TableName() string {
	return "plan"
}
