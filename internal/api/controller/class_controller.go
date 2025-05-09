package controller

import (
	"frs-planning-backend/internal/api/service"
	"frs-planning-backend/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ClassController interface {
	CreateClass(c *gin.Context)
	GetAllClasses(c *gin.Context)
	GetClassByID(c *gin.Context)
	UpdateClass(c *gin.Context)
	DeleteClass(c *gin.Context)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.classService.CreateClass(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Class created successfully"})
}

func (ctrl *classController) GetAllClasses(c *gin.Context) {
	classes, err := ctrl.classService.GetAllClasses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (ctrl *classController) GetClassByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	class, err := ctrl.classService.GetClassByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if class == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Class not found"})
		return
	}

	c.JSON(http.StatusOK, class)
}

func (ctrl *classController) UpdateClass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var request dto.UpdateClassRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.classService.UpdateClass(id, &request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class updated successfully"})
}

func (ctrl *classController) DeleteClass(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	if err := ctrl.classService.DeleteClass(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Class deleted successfully"})
}
