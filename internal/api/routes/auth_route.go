package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Auth(app *gin.Engine, authcontroller controller.AuthController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/auth")
	{
		routes.POST("/login", authcontroller.Login)
		routes.POST("/register", authcontroller.Register)
		routes.GET("/me", middleware.Authenticate(), authcontroller.Me)
	}
}
