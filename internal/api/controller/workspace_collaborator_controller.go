package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	myerror "frs-planning-backend/internal/pkg/error"
	"frs-planning-backend/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type (
	WorkspaceCollaboratorController interface {
		AddCollaborator(ctx *gin.Context)
		GetAllCollaborator(ctx *gin.Context)
		DeleteCollaborator(ctx *gin.Context)
	}

	workspaceCollaboratorController struct {
		workspaceCollaboratorService service.WorskspaceCollaboratorService
	}
)

func NewWorkspaceCOllaborator(workspaceCollaboratorService service.WorskspaceCollaboratorService) WorkspaceCollaboratorController {
	return &workspaceCollaboratorController{
		workspaceCollaboratorService: workspaceCollaboratorService,
	}
}

func (c *workspaceCollaboratorController) AddCollaborator(ctx *gin.Context) {
	var req dto.AddCollaboratorRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	collab, err := c.workspaceCollaboratorService.Add(ctx, req)
	if err != nil {
		response.NewFailed("Failed to add collaborator", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Added collaborator", collab).Send(ctx)
}

func (c *workspaceCollaboratorController) GetAllCollaborator(ctx *gin.Context) {
	workspaceidStr := ctx.Param("workspaceid")

	workspaceid, err := uuid.Parse(workspaceidStr)
	if err != nil {
		response.NewFailed("Invalid workspace ID", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	collab, err := c.workspaceCollaboratorService.Get(ctx, workspaceid)
	if err != nil {
		response.NewFailed("Failed to get collaborators", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Success to get all collaborators", collab).Send(ctx)
}

func (c *workspaceCollaboratorController) DeleteCollaborator(ctx *gin.Context) {
	var req dto.DeleteCollaboratorRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	collab, err := c.workspaceCollaboratorService.Delete(ctx, req)
	if err != nil {
		response.NewFailed("Failed to delete collaborator", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Deleted collaborator", collab).Send(ctx)
}
