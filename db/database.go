package db

import (
	"fmt"
	mylog "frs-planning-backend/internal/pkg/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)

	fmt.Println(mylog.ColorizeInfo("\n=========== Setup Database ==========="))
	mylog.Infof("Connecting to database...")
	db, err := gorm.Open(postgres.Open(DBDSN), &gorm.Config{})
	if err != nil {
		mylog.Errorf("Failed connect to database")
		mylog.Panicf("Failed connect to database")
	}

	mylog.Infof("Success connect to database\n")
	return db
}
