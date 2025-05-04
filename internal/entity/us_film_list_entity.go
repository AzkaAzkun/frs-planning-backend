package entity

import "github.com/google/uuid"

type ListStatus string

const (
	ListStatusPlanToWatch ListStatus = "plan_to_watch"
	ListStatusWatching    ListStatus = "watching"
	ListStatusCompleted   ListStatus = "completed"
	ListStatusOnHold      ListStatus = "on_hold"
	ListStatusDropped     ListStatus = "dropped"
)

type Visibility string

const (
	VisibilityPublic  Visibility = "public"
	VisibilityPrivate Visibility = "private"
)

type FilmList struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FilmId uuid.UUID `json:"film_id"`
	UserId uuid.UUID `json:"user_id"`

	ListStatus ListStatus `json:"list_status"`
	Visibility Visibility `json:"visibility"`

	Film *Film `json:"film" gorm:"foreignKey:FilmId"`

	Timestamp
}

func (f *FilmList) TableName() string {
	return "us_film_lists"
}
