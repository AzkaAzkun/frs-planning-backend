package entity

import "github.com/google/uuid"

type FilmGenre struct {
	ID      uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FilmId  uuid.UUID `json:"film_id"`
	GenreId uuid.UUID `json:"genre_id"`

	Genre *Genre `gorm:"foreignKey:GenreId"`

	Timestamp
}

func (f *FilmGenre) TableName() string {
	return "film_genres"
}
