package entity

import "github.com/google/uuid"

type ReactionStatus string

const (
	StatusLike    ReactionStatus = "like"
	StatusDislike ReactionStatus = "dislike"
)

type Reaction struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ReviewId uuid.UUID `json:"review_id"`
	UserId   uuid.UUID `json:"user_id"`

	Status ReactionStatus `json:"status"`

	Timestamp
}

func (u *Reaction) TableName() string {
	return "us_reaction"
}
