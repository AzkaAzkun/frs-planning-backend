package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Class(router *gin.Engine, classController controller.ClassController, middleware middleware.Middleware) {
	classRoutes := router.Group("/api/v1/classes")
	{
		classRoutes.POST("", middleware.Authenticate(), classController.CreateClass)
		classRoutes.GET("", middleware.Authenticate(), classController.GetAllClasses)
		classRoutes.GET("/:id", middleware.Authenticate(), classController.GetClassByID)
		classRoutes.PUT("/:id", middleware.Authenticate(), classController.UpdateClass)
		classRoutes.DELETE("/:id", middleware.Authenticate(), classController.DeleteClass)
		classRoutes.GET("/course/:course_id", middleware.Authenticate(), classController.GetClassesByCourseID)
	}
}
