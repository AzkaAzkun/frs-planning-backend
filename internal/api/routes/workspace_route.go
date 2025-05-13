package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Workspace(app *gin.Engine, workspaceController controller.WorkspaceController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/workspace")
	{
		routes.POST("/create", workspaceController.CreateWorkspace)
		routes.GET("/:workspaceid", workspaceController.FindWorkspace)
		routes.PUT("/update", workspaceController.UpdateWorkspace)
		routes.DELETE("/delete", workspaceController.DeleteWorkspace)
	}

}
