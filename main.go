package main

import (
	"cat-social-be/app"
	controller "cat-social-be/controller/user"
	"cat-social-be/middleware"
	repository "cat-social-be/repository/user"
	service "cat-social-be/service/user"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := app.NewRouter(userController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	// err := server.ListenAndServe()
	fmt.Print("Server is running on port 3000")
	server.ListenAndServe()
	// helper.PanicIfError(err)
}
