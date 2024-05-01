package app

import (
	controller "cat-social-be/controller/user"

	"github.com/julienschmidt/httprouter"
)

// func NewRouter(catController controller.CatController, userController controller.UserController) *httprouter.Router {
func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	//cat
	// router.GET("/api/cats", catController.FindAll)
	// // router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/cat", catController.Create)
	// router.PUT("/api/cat/:id", catController.Update)
	// router.DELETE("/api/cat/:id", catController.Delete)

	//user
	router.POST("/api/v1/login", userController.Login)
	router.POST("/api/v1/register", userController.Register)

	// router.PanicHandler = exception.ErrorHandler

	return router
}
