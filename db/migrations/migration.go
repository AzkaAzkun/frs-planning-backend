package migrations

import (
	"fmt"
	"frs-planning-backend/internal/entity"
	mylog "frs-planning-backend/internal/pkg/logger"

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
		entity.Course{},
		entity.Class{},
		entity.Workspace{},
		entity.WorkspaceCollaborator{},
		entity.Plan{},
		entity.PlanSettings{},
		entity.ClassSettings{},
	); err != nil {
		return err
	}

	return nil
}
