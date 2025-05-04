package main

import (
	"film-management-api-golang/cmd"
	"film-management-api-golang/internal/config"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to loading env file")
	}

	if err := cmd.Commands(); err != nil {
		panic("Failed Get Command: " + err.Error())
	}

	RestApi := config.NewRest()
	RestApi.Start()
}
