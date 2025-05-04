package controller

import (
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/dto"
	myerror "film-management-api-golang/internal/pkg/error"
	"film-management-api-golang/internal/pkg/response"
	"film-management-api-golang/internal/utils"

	"github.com/gin-gonic/gin"
)

type (
	FilmListController interface {
		Create(ctx *gin.Context)
		UpdateVisibility(ctx *gin.Context)
	}

	filmListController struct {
		filmlistService service.FilmListService
	}
)

func NewFilmList(filmlistService service.FilmListService) FilmListController {
	return &filmListController{
		filmlistService: filmlistService,
	}
}

func (c *filmListController) Create(ctx *gin.Context) {
	var req dto.FilmListRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("invalid input data", myerror.ErrBodyRequest).Send(ctx)
		return
	}

	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		response.NewFailed("failed get user id", err).Send(ctx)
		return
	}

	if err := c.filmlistService.Create(ctx, req, userId); err != nil {
		response.NewFailed("failed create film list", err).Send(ctx)
		return
	}

	response.NewSuccess("success create film list", nil).Send(ctx)
}

func (c *filmListController) UpdateVisibility(ctx *gin.Context) {
	var req dto.FilmListVisibilityRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("invalid input data", myerror.ErrBodyRequest).Send(ctx)
		return
	}

	filmlistId := ctx.Param("id")
	if err := c.filmlistService.UpdateVisibility(ctx, req, filmlistId); err != nil {
		response.NewFailed("failed update film list", err).Send(ctx)
		return
	}

	response.NewSuccess("success update film list", nil).Send(ctx)
}
