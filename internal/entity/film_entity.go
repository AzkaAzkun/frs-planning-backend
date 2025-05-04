package entity

import (
	"time"

	"github.com/google/uuid"
)

type AiringStatus string

const (
	NotYetAired    AiringStatus = "not_yet_aired"
	Airing         AiringStatus = "airing"
	FinishedAiring AiringStatus = "finished_airing"
)

type Film struct {
	ID            uuid.UUID    `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Title         string       `json:"title"`
	Synopsis      string       `json:"synopsis"`
	AiringStatus  AiringStatus `json:"airing_status"`
	TotalEpisodes int          `json:"total_episodes"`
	ReleaseDate   time.Time    `json:"release_date"`

	Images  []FilmImage `json:"images" gorm:"foreignKey:FilmId"`
	Genres  []FilmGenre `json:"genres" gorm:"foreignKey:FilmId"`
	Reviews []Review    `json:"reviews"  gorm:"foreignKey:FilmId"`

	Timestamp
}

func (f *Film) TableName() string {
	return "films"
}
