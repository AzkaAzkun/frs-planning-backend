package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func FilmList(app *gin.Engine, filmlistcontroller controller.FilmListController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/film-lists")
	{
		routes.POST("", middleware.Authenticate(), filmlistcontroller.Create)
		routes.PATCH("/:id", middleware.Authenticate(), filmlistcontroller.UpdateVisibility)
	}
}
