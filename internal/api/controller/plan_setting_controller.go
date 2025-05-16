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
	PlanSettingController interface {
		CreatePlanSetting(c *gin.Context)
		DeletePlanSetting(c *gin.Context)
	}

	planSettingController struct {
		planSettingService service.PlanSettingService
	}
)

func NewPlanSettingController(planSettingService service.PlanSettingService) PlanSettingController {
	return &planSettingController{
		planSettingService: planSettingService,
	}
}

func (c *planSettingController) CreatePlanSetting(ctx *gin.Context) {
	var req dto.PlanSettingCreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	if err := c.planSettingService.Create(ctx.Request.Context(), req); err != nil {
		response.NewFailed("Failed to create plan setting", err).Send(ctx)
		return
	}

	response.NewSuccess("Created a plan setting", nil).Send(ctx)
}

func (c *planSettingController) DeletePlanSetting(ctx *gin.Context) {
	planSettingId := ctx.Param("id")
	if err := c.planSettingService.Delete(ctx.Request.Context(), planSettingId); err != nil {
		response.NewFailed("Failed to delete plan setting", err).Send(ctx)
		return
	}

	response.NewSuccess("Deleted a plan setting", nil).Send(ctx)
}
