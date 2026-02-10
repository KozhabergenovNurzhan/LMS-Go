package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/KozhabergenovNurzhan/GoProj1/internal/models"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	courseService *service.CourseService
	// TODO add other service
}

func (h *Handler) InitRoutes() (*gin.Engine, error) {
	r := gin.New()

	r.GET("/courses", h.GetCourses)
	r.GET("/courses/:id", h.GetCourseById)
	r.DELETE("/courses/:id", h.DeleteCourse)
	r.POST("/courses", h.CreateCourse)

	return r, nil
}

func NewHandler(cs *service.CourseService) *Handler {
	return &Handler{
		courseService: cs,
	}
}

func (h *Handler) DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
		return
	}

	err = h.courseService.DeleteByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "course to delete not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{
		"message": "course is successfully deleted",
	})
}

func (h *Handler) GetCourses(c *gin.Context) {
	courses, err := h.courseService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to select data",
		})
		return
	}

	c.JSON(http.StatusOK, courses)
}

func (h *Handler) GetCourseById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course id"})
		return
	}

	course, err := h.courseService.GetByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, course)
}

func (h *Handler) CreateCourse(c *gin.Context) {
	var input models.CreateCourse

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	id, err := h.courseService.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create",
		})
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}
