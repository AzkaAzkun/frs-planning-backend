package controller

import (
	"film-management-api-golang/internal/api/service"
	"film-management-api-golang/internal/dto"
	myerror "film-management-api-golang/internal/pkg/error"
	"film-management-api-golang/internal/pkg/meta"
	"film-management-api-golang/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	FilmController interface {
		Create(ctx *gin.Context)
		GetListFilm(ctx *gin.Context)
		GetDetailFilm(ctx *gin.Context)
	}

	filmController struct {
		filmService service.FilmService
	}
)

func NewFilm(filmService service.FilmService) FilmController {
	return &filmController{
		filmService: filmService,
	}
}

func (c *filmController) Create(ctx *gin.Context) {
	var req dto.FilmCreateRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get input data", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	form, _ := ctx.MultipartForm()
	req.Images = form.File["images"]

	result, err := c.filmService.Create(ctx, req)
	if err != nil {
		response.NewFailed("failed create new film", err).Send(ctx)
		return
	}

	response.NewSuccess("success create new film", result).Send(ctx)
}

func (c *filmController) GetListFilm(ctx *gin.Context) {
	result, err := c.filmService.GetListFilm(ctx, meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get list film", err).Send(ctx)
		return
	}

	response.NewSuccess("success get list film", result.Data, result.Meta).Send(ctx)
}

func (c *filmController) GetDetailFilm(ctx *gin.Context) {
	filmId := ctx.Param("id")
	result, err := c.filmService.GetDetailFilm(ctx, filmId)
	if err != nil {
		response.NewFailed("failed get detail film", err).Send(ctx)
		return
	}

	response.NewSuccess("success get detail film", result).Send(ctx)
}
