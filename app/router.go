package app

import (
	catController "cat-social-be/controller/cat"
	matchController "cat-social-be/controller/match"
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
	catMiddlewareRoutes := r.Group("/v1/cat")
	catMiddlewareRoutes.Use(middleware.JwtAuthMiddleware())
	catMiddlewareRoutes.GET("/", catController.GetCats)
	catMiddlewareRoutes.POST("/", catController.CreateCat)
	catMiddlewareRoutes.PUT("/:id", catController.UpdateCat)
	catMiddlewareRoutes.DELETE("/:id", catController.DeleteCat)
	catMiddlewareRoutes.POST("/match", matchController.CreateMatch)
	catMiddlewareRoutes.GET("/match", matchController.GetMatches)
	catMiddlewareRoutes.POST("/match/approve", matchController.ApproveMatch)
	catMiddlewareRoutes.POST("/match/reject", matchController.RejectMatch)
	catMiddlewareRoutes.DELETE("/match/:id", matchController.DeleteMatch)
	//user
	r.POST("/v1/user/login", userController.Login)
	r.POST("/v1/user/register", userController.Register)

	//match

	// router.PanicHandler = exception.ErrorHandler

	return r
}
