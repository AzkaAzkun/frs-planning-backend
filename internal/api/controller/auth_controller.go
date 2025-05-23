package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	myerror "frs-planning-backend/internal/pkg/error"
	"frs-planning-backend/internal/pkg/response"
	"frs-planning-backend/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	AuthController interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
		Verify(ctx *gin.Context)
		ForgotPassword(ctx *gin.Context)
		ChangePassword(ctx *gin.Context)
		Me(ctx *gin.Context)
	}

	authController struct {
		authService service.AuthService
	}
)

func NewAuth(authService service.AuthService) AuthController {
	return &authController{
		authService: authService,
	}
}

func (c *authController) Register(ctx *gin.Context) {
	var req dto.RegisterRequest

	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", myerror.New(err.Error(), http.StatusBadRequest)).Send(ctx)
		return
	}

	user, err := c.authService.Register(ctx, req)
	if err != nil {
		response.NewFailed("failed register account", err).Send(ctx)
		return
	}

	response.NewSuccess("success register account", user).Send(ctx)
}

func (c *authController) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBind(&req); err != nil {
		response.NewFailed("failed get data from body", err).Send(ctx)
		return
	}

	result, err := c.authService.Login(ctx.Request.Context(), req)
	if err != nil {
		response.NewFailed("failed login", err).Send(ctx)
		return
	}

	response.NewSuccess("success login", result).Send(ctx)
}

func (c *authController) Verify(ctx *gin.Context) {
	token := ctx.Query("token")
	if err := c.authService.Verify(ctx, token); err != nil {
		response.NewFailed("failed verify account", err).Send(ctx)
		return
	}

	response.NewSuccess("success verify account", nil).Send(ctx)
}

func (c *authController) ForgotPassword(ctx *gin.Context) {

}

func (c *authController) ChangePassword(ctx *gin.Context) {

}

func (c *authController) Me(ctx *gin.Context) {
	userId, err := utils.GetUserIdFromCtx(ctx)
	if err != nil {
		response.NewFailed("failed get user id", err).Send(ctx)
		return
	}

	res, err := c.authService.GetMe(ctx.Request.Context(), userId)
	if err != nil {
		response.NewFailed("failed get me", err).Send(ctx)
		return
	}

	response.NewSuccess("success get me", res).Send(ctx)
}
