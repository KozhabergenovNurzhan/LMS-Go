package handler

import (
	"net/http"

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
	// TODO implement other endpoints

	return r, nil
}

func NewHandler(cs *service.CourseService) *Handler {
	return &Handler{
		courseService: cs,
	}
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
