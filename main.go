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
	// validate := validator.New()
	// userRepository := repository.NewUserRepository()
	// userService := service.NewUserService(userRepository, db, validate)
	// userController := controller.NewUserController(userService)
	// router := app.NewRouter(userController)

	// server := http.Server{
	// 	Addr:    "localhost:3000",
	// 	Handler: router,
	// }

	// // err := server.ListenAndServe()
	// fmt.Print("Server is running on port 3000")
	// err := server.ListenAndServe()
	// helper.PanicIfError(err)
	r := app.SetupRouter(db)
	r.Run()
}
