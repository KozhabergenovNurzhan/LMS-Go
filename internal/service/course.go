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

func (cs *CourseService) GetByID(id int) (models.Course, error) {
	return cs.repo.GetByID(id)
}

func (cs *CourseService) DeleteByID(id int) error {
	return cs.repo.DeleteByID(id)
}

func (cs *CourseService) Create(input models.CreateCourse) (int, error) {
	return cs.repo.Create(input)
}
