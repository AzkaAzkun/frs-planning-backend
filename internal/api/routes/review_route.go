package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Review(app *gin.Engine, reviewcontroller controller.ReviewController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/reviews")
	{
		routes.POST("", middleware.Authenticate(), reviewcontroller.Create)
	}
}
