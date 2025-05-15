package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	myerror "frs-planning-backend/internal/pkg/error"
	"frs-planning-backend/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	ClassSettingController interface {
		Create(ctx *gin.Context)
		Clone(ctx *gin.Context)
	}

	classSettingController struct {
		classSettingService service.ClassSettingService
	}
)

func NewClassSettingController(classsettingservice service.ClassSettingService) ClassSettingController {
	return &classSettingController{
		classSettingService: classsettingservice,
	}
}

func (c *classSettingController) Create(ctx *gin.Context) {
	UserId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uidStr, ok := UserId.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	var req dto.CreateClassSettingRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	classSetting, err := c.classSettingService.Create(ctx, req, uidStr)
	if err != nil {
		response.NewFailed("Failed to create new class setting", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Created new class setting", classSetting).Send(ctx)
}

func (c *classSettingController) Clone(ctx *gin.Context) {
	UserId, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	uidStr, ok := UserId.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	var req dto.CloneClassSettingRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	cloneClassSetting, err := c.classSettingService.Clone(ctx, uidStr, req)
	if err != nil {
		response.NewFailed("Failed to clone class setting", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}
	response.NewSuccess("Cloned class setting", cloneClassSetting)
}
