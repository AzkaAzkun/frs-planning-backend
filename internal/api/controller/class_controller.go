package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type ClassController interface {
	CreateClass(c *gin.Context)
	GetAllClasses(c *gin.Context)
	GetClassByID(c *gin.Context)
	UpdateClass(c *gin.Context)
	DeleteClass(c *gin.Context)
	GetClassesByCourseID(c *gin.Context)
}

type classController struct {
	classService service.ClassService
}

func NewClassController(classService service.ClassService) ClassController {
	return &classController{
		classService: classService,
	}
}

func (ctrl *classController) CreateClass(c *gin.Context) {
	var request dto.CreateClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewFailed("failed get data from request", err).Send(c)
		return
	}

	res, err := ctrl.classService.CreateClass(c.Request.Context(), request)
	if err != nil {
		response.NewFailed("failed create class", err).Send(c)
		return
	}

	response.NewSuccess("success create class", res).Send(c)
}

func (ctrl *classController) GetAllClasses(c *gin.Context) {
	classes, err := ctrl.classService.GetAllClasses(c.Request.Context())
	if err != nil {
		response.NewFailed("failed get all classes", err).Send(c)
		return
	}

	response.NewSuccess("success get all classes", classes).Send(c)
}

func (ctrl *classController) GetClassByID(c *gin.Context) {
	id := c.Param("id")
	class, err := ctrl.classService.GetClassByID(c.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get class by id", err).Send(c)
		return
	}

	response.NewSuccess("success get class by id", class).Send(c)
}

func (ctrl *classController) UpdateClass(c *gin.Context) {
	id := c.Param("id")

	var request dto.UpdateClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewFailed("failed get data from request", err).Send(c)
		return
	}

	if err := ctrl.classService.UpdateClass(c.Request.Context(), id, request); err != nil {
		response.NewFailed("failed update class by id", err).Send(c)
		return
	}

	response.NewSuccess("success update class", nil).Send(c)
}

func (ctrl *classController) DeleteClass(c *gin.Context) {
	id := c.Param("id")
	if err := ctrl.classService.DeleteClass(c.Request.Context(), id); err != nil {
		response.NewFailed("failed delete class", err).Send(c)

		return
	}

	response.NewSuccess("success delete class", nil).Send(c)
}

func (ctrl *classController) GetClassesByCourseID(c *gin.Context) {
	courseID := c.Param("course_id")

	classes, err := ctrl.classService.GetClassesByCourseID(c.Request.Context(), courseID)
	if err != nil {
		response.NewFailed("failed get classes by course id", err).Send(c)
		return
	}

	response.NewSuccess("success get classes by course id", classes).Send(c)
}
