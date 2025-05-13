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
	WorkspaceController interface {
		CreateWorkspace(ctx *gin.Context)
		FindWorkspace(ctx *gin.Context)
		UpdateWorkspace(ctx *gin.Context)
		DeleteWorkspace(ctx *gin.Context)
	}
	workspaceController struct {
		workspaceService service.WorkspaceService
	}
)

func NewWorkspace(workspaceService service.WorkspaceService) WorkspaceController {
	return &workspaceController{
		workspaceService: workspaceService,
	}
}

func (c *workspaceController) CreateWorkspace(ctx *gin.Context) {
	var req dto.CreateWorkspaceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	workspace, err := c.workspaceService.Create(ctx, req)
	if err != nil {
		response.NewFailed("Failed to create workspace", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Created a workspace", workspace).Send(ctx)
}

func (c *workspaceController) FindWorkspace(ctx *gin.Context) {
	workspaceid := ctx.Param("workspaceid")

	workspaceUUID, err := uuid.Parse(workspaceid)
	if err != nil {
		response.NewFailed("Invalid workspace ID", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	workspace, err := c.workspaceService.Find(ctx, workspaceUUID)
	if err != nil {
		response.NewFailed("Failed to find workspace", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	response.NewSuccess("Found workspace", workspace).Send(ctx)
}

func (c *workspaceController) UpdateWorkspace(ctx *gin.Context) {
	var req dto.UpdateWorkspaceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	workspace, err := c.workspaceService.Update(ctx, req)
	if err != nil {
		response.NewFailed("Failed to update workspace", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Updated workspace", workspace).Send(ctx)
}

func (c *workspaceController) DeleteWorkspace(ctx *gin.Context) {
	var req dto.DeleteWorkspaceRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	workspace, err := c.workspaceService.Delete(ctx, req)
	if err != nil {
		response.NewFailed("Failed to delete workspace", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Deleted workspace", workspace).Send(ctx)
}
