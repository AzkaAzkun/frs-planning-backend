package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Reaction(app *gin.Engine, reactioncontroller controller.ReactionController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/reactions")
	{
		routes.POST("", middleware.Authenticate(), reactioncontroller.Create)
		routes.PUT("/:id", middleware.Authenticate(), reactioncontroller.Update)
	}
}
