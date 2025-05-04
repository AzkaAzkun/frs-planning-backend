package seeders

import (
	"film-management-api-golang/db/seeder/seeds"
	mylog "film-management-api-golang/internal/pkg/logger"
	"fmt"

	"gorm.io/gorm"
)

func Seeding(db *gorm.DB) error {
	seeders := []func(*gorm.DB) error{
		seeds.SeederUser,
		seeds.SeederGenre,
	}

	fmt.Println(mylog.ColorizeInfo("\n=========== Start Seeding ==========="))
	for _, seeder := range seeders {
		if err := seeder(db); err != nil {
			return err
		}
	}

	return nil
}
