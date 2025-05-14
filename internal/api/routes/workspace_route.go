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
		routes.GET("/:workspaceid", middleware.Authenticate(), workspaceController.FindWorkspace)
		routes.PUT("/update", middleware.Authenticate(), workspaceController.UpdateWorkspace)
		routes.DELETE("/delete", middleware.Authenticate(), workspaceController.DeleteWorkspace)
	}

}
