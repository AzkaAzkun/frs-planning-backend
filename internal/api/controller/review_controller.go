package controller

import (
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/pkg/response"
	"film-management-api-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type (
	ReviewController interface {
		Create(ctx *gin.Context)
	}

	reviewController struct {
		reviewService service.ReviewService
	}
)

func NewReview(reviewService service.ReviewService) ReviewController {
	return &reviewController{
		reviewService: reviewService,
	}
}

func (c *reviewController) Create(ctx *gin.Context) {
	var req dto.ReviewRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewSuccess("invalid input data", err).Send(ctx)
		return
	}

	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		response.NewFailed("failed get user id", err).Send(ctx)
		return
	}

	if err := c.reviewService.Create(ctx, req, userId); err != nil {
		response.NewFailed("failed create review", err).Send(ctx)
		return
	}

	response.NewSuccess("success create review", nil).Send(ctx)
}
