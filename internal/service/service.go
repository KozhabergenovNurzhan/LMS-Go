package service

import (
	"github.com/KozhabergenovNurzhan/GoProj1/internal/models"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/repository"
)

type CourseService struct {
	repo repository.CourseRepo
}

func NewCourseService(repo repository.CourseRepo) *CourseService {
	return &CourseService{repo: repo}
}

func (cs *CourseService) GetAll() ([]models.Course, error) {
	return cs.repo.GetAll()
}
