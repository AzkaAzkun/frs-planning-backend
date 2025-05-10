package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.courseService.CreateCourse(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Course created successfully"})
}

func (ctrl *courseController) GetAllCourses(c *gin.Context) {
	courses, err := ctrl.courseService.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (ctrl *courseController) GetCourseByID(c *gin.Context) {
	id := c.Param("id")

	course, err := ctrl.courseService.GetCourseByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if course == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (ctrl *courseController) UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var request dto.UpdateCourseRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.courseService.UpdateCourse(id, &request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

func (ctrl *courseController) DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.courseService.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
