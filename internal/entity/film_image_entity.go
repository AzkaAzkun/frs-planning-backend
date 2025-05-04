package entity

import "github.com/google/uuid"

type FilmImage struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FilmId uuid.UUID `json:"film_id"`

	ImagePath string `json:"image_path"`

	Timestamp
}

func (f *FilmImage) TableName() string {
	return "film_images"
}
