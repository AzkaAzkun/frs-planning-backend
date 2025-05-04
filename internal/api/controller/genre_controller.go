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
	GenreController interface {
		Create(ctx *gin.Context)
		GetAll(ctx *gin.Context)
		GetAllPaginated(ctx *gin.Context)
		Update(ctx *gin.Context)
	}

	genreController struct {
		genreService service.GenreService
	}
)

func NewGenre(genreService service.GenreService) GenreController {
	return &genreController{
		genreService: genreService,
	}
}

func (c *genreController) Create(ctx *gin.Context) {
	var req dto.GenreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("failed create genre", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	genre, err := c.genreService.Create(ctx, req)
	if err != nil {
		response.NewFailed("failed create genre", err).Send(ctx)
		return
	}

	response.NewSuccess("success create genre", genre).Send(ctx)
}

func (c *genreController) GetAll(ctx *gin.Context) {
	result, err := c.genreService.GetAll(ctx)
	if err != nil {
		response.NewFailed("failed get all genre", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all genre", result).Send(ctx)
}

func (c *genreController) GetAllPaginated(ctx *gin.Context) {
	result, err := c.genreService.GetAllPaginated(ctx, meta.New(ctx))
	if err != nil {
		response.NewFailed("failed get all genre", err).Send(ctx)
		return
	}

	response.NewSuccess("success get all genre", result.Data, result.Meta).Send(ctx)
}

func (c *genreController) Update(ctx *gin.Context) {
	var req dto.GenreRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.NewFailed("failed update genre", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	id := ctx.Param("id")
	genre, err := c.genreService.Update(ctx, req, id)
	if err != nil {
		response.NewFailed("failed update genre", err).Send(ctx)
		return
	}

	response.NewSuccess("success update genre", genre).Send(ctx)
}
