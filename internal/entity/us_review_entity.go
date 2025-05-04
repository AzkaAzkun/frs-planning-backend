package entity

import "github.com/google/uuid"

type Review struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FilmId uuid.UUID `json:"film_id"`
	UserId uuid.UUID `json:"user_id"`

	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`

	Film *Film `gorm:"foreignKey:FilmId"`

	Timestamp
}

func (u *Review) TableName() string {
	return "us_reviews"
}
