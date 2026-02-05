package repository

import "github.com/KozhabergenovNurzhan/GoProj1/internal/models"

type CourseRepo interface {
	GetAll() ([]models.Course, error)
	// TODO implement other methods
}

type PsgCourseRepo struct {
	db *DB
}

func NewPsgCourseRepo(db *DB) *PsgCourseRepo {
	return &PsgCourseRepo{
		db: db,
	}
}

func (pcr *PsgCourseRepo) GetAll() ([]models.Course, error) {
	// TODO implement use of DB
	return []models.Course{
		{ID: 1, Name: "Python course"},
		{ID: 2, Name: "Go course"},
		{ID: 3, Name: "Java course"},
	}, nil
}
