package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func User(app *gin.Engine, usercontroller controller.UserController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/users")
	{
		routes.GET("/:id", usercontroller.GetById)
	}
}
