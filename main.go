package main

import (
	"MovieApi/app"
	"MovieApi/config"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()

	err := config.InitPostgres()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
