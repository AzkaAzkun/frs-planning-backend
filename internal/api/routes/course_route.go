package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/api/repository"
	"frs-planning-backend/internal/api/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCourseRoutes(router *gin.Engine, db *gorm.DB) {
	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseController := controller.NewCourseController(courseService)

	courseRoutes := router.Group("/api/courses")
	{
		courseRoutes.POST("", courseController.CreateCourse)
		courseRoutes.GET("", courseController.GetAllCourses)
		courseRoutes.GET("/:id", courseController.GetCourseByID)
		courseRoutes.PUT("/:id", courseController.UpdateCourse)
		courseRoutes.DELETE("/:id", courseController.DeleteCourse)
	}
}
