package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Plan(router *gin.Engine, planController controller.PlanController, middleware middleware.Middleware) {
	planRoutes := router.Group("/api/v1/plans")
	{
		planRoutes.POST("", middleware.Authenticate(), planController.CreatePlan)
		planRoutes.GET(":id/workspaces", middleware.Authenticate(), planController.GetAllPlans)
		// planRoutes.GET("/:id", planController.GetCourseByID)
		planRoutes.PUT("/:id", middleware.Authenticate(), planController.UpdatePlan)
		planRoutes.DELETE("/:id", middleware.Authenticate(), planController.DeletePlan)
	}
}
