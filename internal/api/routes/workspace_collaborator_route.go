package routes

import (
	"frs-planning-backend/internal/api/controller"
	"frs-planning-backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func WorkspaceCollaborator(app *gin.Engine, workspacecollaborartor controller.WorkspaceCollaboratorController, middleware middleware.Middleware) {
	routes := app.Group("/api/v1/workspace")
	{
		routes.POST("/add", middleware.Authenticate(), workspacecollaborartor.AddCollaborator)
		routes.GET("/collaborators/:workspaceid", middleware.Authenticate(), workspacecollaborartor.GetAllCollaborator)
		routes.DELETE("/remove", middleware.Authenticate(), workspacecollaborartor.DeleteCollaborator)
	}
}
