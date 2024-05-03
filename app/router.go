package app

import (
	catController "cat-social-be/controller/cat"
	userController "cat-social-be/controller/user"
	"cat-social-be/middleware"
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
	catMiddlewareRoutes := r.Group("/api/v1/cats")
	catMiddlewareRoutes.Use(middleware.JwtAuthMiddleware())
	catMiddlewareRoutes.GET("/", catController.GetCats)
	catMiddlewareRoutes.POST("/", catController.CreateCat)
	catMiddlewareRoutes.PUT("/:id", catController.UpdateCat)
	catMiddlewareRoutes.DELETE("/:id", catController.DeleteCat)
	//user
	r.POST("/api/v1/login", userController.Login)
	r.POST("/api/v1/register", userController.Register)
	// router.PanicHandler = exception.ErrorHandler

	return r
}
