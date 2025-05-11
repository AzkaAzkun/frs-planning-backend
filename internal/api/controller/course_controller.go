package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	"frs-planning-backend/internal/pkg/response"

	"github.com/gin-gonic/gin"
)

type CourseController interface {
	CreateCourse(c *gin.Context)
	GetAllCourses(c *gin.Context)
	GetCourseByID(c *gin.Context)
	UpdateCourse(c *gin.Context)
	DeleteCourse(c *gin.Context)
}

type courseController struct {
	courseService service.CourseService
}

func NewCourseController(courseService service.CourseService) CourseController {
	return &courseController{
		courseService: courseService,
	}
}

func (ctrl *courseController) CreateCourse(c *gin.Context) {
	var request dto.CreateCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewFailed("failed get data from request", err).Send(c)
		return
	}

	res, err := ctrl.courseService.CreateCourse(c.Request.Context(), request)
	if err != nil {
		response.NewFailed("failed create course", err).Send(c)
		return
	}

	response.NewSuccess("success create course", res).Send(c)
}

func (ctrl *courseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.courseService.GetAllCourses(c.Request.Context())
	if err != nil {
		response.NewFailed("failed get all courses", err).Send(c)
		return
	}

	response.NewSuccess("success get all courses", courses).Send(c)
}

func (ctrl *courseController) GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	course, err := ctrl.courseService.GetCourseByID(c.Request.Context(), id)
	if err != nil {
		response.NewFailed("failed get course by id", err).Send(c)
		return
	}

	response.NewSuccess("success get course by id", course).Send(c)
}

func (ctrl *courseController) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var request dto.UpdateCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.NewFailed("failed get data from request", err).Send(c)
		return
	}

	if err := ctrl.courseService.UpdateCourse(c.Request.Context(), id, request); err != nil {
		response.NewFailed("failed update course by id", err).Send(c)
		return
	}

	response.NewSuccess("success update course by id", nil).Send(c)
}

func (ctrl *courseController) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.courseService.DeleteCourse(c.Request.Context(), id); err != nil {
		response.NewFailed("failed delete course by id", err).Send(c)
		return
	}

	response.NewSuccess("success delete course by id", nil).Send(c)
}
