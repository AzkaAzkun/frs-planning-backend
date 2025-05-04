package migrations

import (
	"film-management-api-golang/internal/entity"
	mylog "film-management-api-golang/internal/pkg/logger"
	"fmt"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	fmt.Println(mylog.ColorizeInfo("\n=========== Start Migrate ==========="))
	mylog.Infof("Migrating Tables...")

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		return err
	}

	//migrate table
	if err := db.AutoMigrate(
		entity.User{},
		entity.Genre{},
		entity.Film{},
		entity.FilmGenre{},
		entity.FilmImage{},
		entity.FilmList{},
		entity.Review{},
		entity.Reaction{},
	); err != nil {
		return err
	}

	return nil
}
