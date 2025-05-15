package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	myerror "frs-planning-backend/internal/pkg/error"
	"frs-planning-backend/internal/pkg/meta"
	"frs-planning-backend/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	PlanController interface {
		CreatePlan(c *gin.Context)
		GetAllPlans(c *gin.Context)
		// GetPlanDetail(c *gin.Context)
		// UpdatePlan(c *gin.Context)
		// DeletePlan(c *gin.Context)
	}

	planController struct {
		planService service.PlanService
	}
)

func NewPlanController(planService service.PlanService) PlanController {
	return &planController{
		planService: planService,
	}
}

func (p *planController) CreatePlan(c *gin.Context) {
	var req dto.PlanCreateRequest
	if err := c.ShouldBind(&req); err != nil {
		response.NewFailed("Failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(c)
		return
	}
	plan, err := p.planService.CreatePlan(c, req)
	if err != nil {
		response.NewFailed("Failed to create plan", err).Send(c)
		return
	}

	response.NewSuccess("Created a plan", plan).Send(c)
}

func (p *planController) GetAllPlans(c *gin.Context) {
	workspaceId := c.Param("id")

	result, err := p.planService.GetAllPlans(c, workspaceId, meta.New(c))
	if err != nil {
		response.NewFailed("Failed to get all plans", err).Send(c)
		return
	}

	response.NewSuccess("Get all plans", result.Plans, result.Meta).Send(c)
}
