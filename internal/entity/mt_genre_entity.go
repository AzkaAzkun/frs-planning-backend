package entity

import "github.com/google/uuid"

type Genre struct {
	ID   uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name"`

	Timestamp
}

func (g *Genre) TableName() string {
	return "mt_genres"
}
