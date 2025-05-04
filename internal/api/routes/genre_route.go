package routes

import (
	"film-management-api-golang/internal/api/controller"
	"film-management-api-golang/internal/entity"
	"film-management-api-golang/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Genre(app *gin.Engine, genrecontroller controller.GenreController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/genres")
	{
		routes.POST("", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), genrecontroller.Create)
		routes.GET("/admin", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), genrecontroller.GetAllPaginated)
		routes.GET("", genrecontroller.GetAll)
		routes.PUT("/:id", middleware.Authenticate(), middleware.OnlyAllow(string(entity.RoleAdmin)), genrecontroller.Update)
	}
}
