package app

import (
	controller "cat-social-be/controller/user"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	//cat
	// router.GET("/api/cats", catController.FindAll)
	// // router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/cat", catController.Create)
	// router.PUT("/api/cat/:id", catController.Update)
	// router.DELETE("/api/cat/:id", catController.Delete)

	//user

	r.POST("/api/v1/login", controller.Login)

	// authMiddlewareRoutes := r.Group("/api/v1/auth")
	// authMiddlewareRoutes.Use(middleware.JwtAuthMiddleware())
	// authMiddlewareRoutes.POST("/register", controller.Register)
	r.POST("/api/v1/auth/register", controller.Register)
	// router.PanicHandler = exception.ErrorHandler

	return r
}
