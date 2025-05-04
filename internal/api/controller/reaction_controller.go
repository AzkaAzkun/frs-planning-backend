package controller

import (
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/dto"
	"film-management-api-golang/internal/pkg/response"
	"film-management-api-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type (
	ReactionController interface {
		Create(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	reactionController struct {
		reactionService service.ReactionService
	}
)

func NewReaction(reactionService service.ReactionService) ReactionController {
	return &reactionController{
		reactionService: reactionService,
	}
}

func (c *reactionController) Create(ctx *gin.Context) {
	var req dto.ReactionRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("invalid input data", err).Send(ctx)
		return
	}

	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		response.NewFailed("failed get user id", err).Send(ctx)
		return
	}

	if err := c.reactionService.Create(ctx.Request.Context(), req, userId); err != nil {
		response.NewFailed("failed create reaction", err).Send(ctx)
		return
	}

	response.NewSuccess("success create reaction", nil).Send(ctx)
}

func (c *reactionController) Update(ctx *gin.Context) {
	var req dto.ReactionUpdate
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewSuccess("invalid input data", err).Send(ctx)
		return
	}

	reactionId := ctx.Param("id")
	if err := c.reactionService.Update(ctx.Request.Context(), req, reactionId); err != nil {
		response.NewFailed("failed update reaction", err).Send(ctx)
		return
	}

	response.NewSuccess("success update reaction", nil).Send(ctx)
}
