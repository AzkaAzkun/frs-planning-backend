package seeds

import (
	"encoding/json"
	"film-management-api-golang/internal/entity"
	mylog "film-management-api-golang/internal/pkg/logger"
	"os"

	"gorm.io/gorm"
)

func SeederUser(db *gorm.DB) error {
	mylog.Infof("Seeding users...")
	jsonFile, err := os.Open("./db/seeder/data/user_data.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	var listEntity []entity.User
	if err := json.NewDecoder(jsonFile).Decode(&listEntity); err != nil {
		return err
	}

	for _, entity := range listEntity {
		if err := db.Save(&entity).Error; err != nil {
			return err
		}
	}

	mylog.Infof("Seeding users completed")
	return nil
}
