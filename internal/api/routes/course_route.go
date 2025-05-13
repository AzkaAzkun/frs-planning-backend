package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Course(router *gin.Engine, courseController controller.CourseController, middleware middleware.Middleware) {
	courseRoutes := router.Group("/api/courses")
	{
		courseRoutes.POST("", courseController.CreateCourse)
		courseRoutes.GET("", courseController.GetAllCourses)
		courseRoutes.GET("/:id", courseController.GetCourseByID)
		courseRoutes.PUT("/:id", courseController.UpdateCourse)
		courseRoutes.DELETE("/:id", courseController.DeleteCourse)
	}
}
