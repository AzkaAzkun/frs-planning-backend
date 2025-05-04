package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Film(app *gin.Engine, filmcontroller controller.FilmController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/films")
	{
		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), filmcontroller.Create)
		routes.GET("", filmcontroller.GetListFilm)
		routes.GET("/:id", filmcontroller.GetDetailFilm)
	}
}
