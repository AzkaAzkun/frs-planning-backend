package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Workspace(app *gin.Engine, workspaceController controller.WorkspaceController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/workspace")
	{
		routes.POST("/create", middleware.Authenticate(), workspaceController.CreateWorkspace)
		routes.GET("/:id", middleware.Authenticate(), workspaceController.FindWorkspace)
		routes.GET("/get", middleware.Authenticate(), workspaceController.GetWorkspace)
		routes.PUT("/update", middleware.Authenticate(), workspaceController.UpdateWorkspace)
		routes.DELETE("/delete/:id", middleware.Authenticate(), workspaceController.DeleteWorkspace)
	}

}
