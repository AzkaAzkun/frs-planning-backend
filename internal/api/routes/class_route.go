package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterClassRoutes(router *gin.Engine, db *gorm.DB) {
	classRepo := repository.NewClassRepository(db)
	classService := service.NewClassService(classRepo)
	classController := controller.NewClassController(classService)

	classRoutes := router.Group("/api/classes")
	{
		classRoutes.POST("", classController.CreateClass)
		classRoutes.GET("", classController.GetAllClasses)
		classRoutes.GET("/:id", classController.GetClassByID)
		classRoutes.PUT("/:id", classController.UpdateClass)
		classRoutes.DELETE("/:id", classController.DeleteClass)
	}
}
