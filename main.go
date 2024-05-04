package main

import (
	"cat-social-be/app"
	"cat-social-be/helper"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	environment := helper.Getenv("ENVIRONMENT", "development")

	if environment == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	db := app.NewDB()
	r := app.SetupRouter(db)
	r.Run(":8080")
}
