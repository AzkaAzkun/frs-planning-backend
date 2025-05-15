package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func ClassSetting(app *gin.Engine, classsettingcontroller controller.ClassSettingController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/class-setting")
	{
		routes.POST("/create", middleware.Authenticate(), classsettingcontroller.Create)
		routes.POST("/clone", middleware.Authenticate(), classsettingcontroller.Clone)
		routes.GET("", classsettingcontroller.GetAll)
		routes.GET("/private", middleware.Authenticate(), classsettingcontroller.GetAllPrivate)
	}
}
