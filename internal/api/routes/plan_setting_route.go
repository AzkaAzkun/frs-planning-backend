package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PlanSetting(router *gin.Engine, planSettingController controller.PlanSettingController, middleware middleware.Middleware) {
	planRoutes := router.Group("/api/v1/plans/settings")
	{
		planRoutes.POST("", middleware.Authenticate(), planSettingController.CreatePlanSetting)
		planRoutes.DELETE("/:id", middleware.Authenticate(), planSettingController.DeletePlanSetting)
	}
}
