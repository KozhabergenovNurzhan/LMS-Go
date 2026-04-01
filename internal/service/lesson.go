package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/KozhabergenovNurzhan/GoProj1/internal/apperror"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/models"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/repository"
	"github.com/jmoiron/sqlx"
)

type LessonService struct {
	repo       repository.LessonRepo
	courseRepo repository.CourseRepo
	db         *sqlx.DB
}

func NewLessonService(repo repository.LessonRepo, courseRepo repository.CourseRepo, db *sqlx.DB) *LessonService {
	return &LessonService{
		repo:       repo,
		courseRepo: courseRepo,
		db:         db,
	}
}

func (ls *LessonService) GetAll() ([]models.Lesson, error) {
	lessons, err := ls.repo.GetAll()
	if err != nil {
		return nil, apperror.New(http.StatusInternalServerError, "failed to get lessons", err)
	}
	return lessons, nil
}

func (ls *LessonService) GetByID(id int) (models.Lesson, error) {
	lesson, err := ls.repo.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrLessonNotFound):
			return models.Lesson{}, apperror.New(http.StatusNotFound, "lesson not found", err)
		default:
			return models.Lesson{}, apperror.New(http.StatusInternalServerError, "failed to get lesson", err)
		}
	}
	return lesson, nil
}

func (ls *LessonService) DeleteByID(ctx context.Context, id int) error {
	lesson, err := ls.repo.GetByID(id)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrLessonNotFound):
			return apperror.New(http.StatusNotFound, "lesson not found", err)
		default:
			return apperror.New(http.StatusInternalServerError, "failed to get lesson", err)
		}
	}

	course, err := ls.courseRepo.GetByID(ctx, lesson.CourseID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrCourseNotFound):
			return apperror.New(http.StatusNotFound, "course not found", err)
		default:
			return apperror.New(http.StatusInternalServerError, "failed to get course", err)
		}
	}

	if course.IsActive {
		return apperror.New(http.StatusConflict, "cannot delete lesson inside active course", nil)
	}

	if err = ls.repo.DeleteByID(ctx, id); err != nil {
		switch {
		case errors.Is(err, models.ErrLessonNotFound):
			return apperror.New(http.StatusNotFound, "lesson not found", err)
		default:
			return apperror.New(http.StatusInternalServerError, "failed to delete lesson", err)
		}
	}
	return nil
}

func (ls *LessonService) Create(ctx context.Context, input models.CreateLesson) (int, error) {
	if err := input.Validate(); err != nil {
		return 0, apperror.New(http.StatusBadRequest, err.Error(), nil)
	}

	_, err := ls.courseRepo.GetByID(ctx, input.CourseID)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrCourseNotFound):
			return 0, apperror.New(http.StatusNotFound, "course not found", err)
		default:
			return 0, apperror.New(http.StatusInternalServerError, "failed to get course", err)
		}
	}

	id, err := ls.repo.Create(ctx, input)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrCourseNotFound):
			return 0, apperror.New(http.StatusNotFound, "course not found", err)
		default:
			return 0, apperror.New(http.StatusInternalServerError, "failed to create lesson", err)
		}
	}
	return id, nil
}

func (ls *LessonService) Update(id int, input models.UpdateLesson) (int, error) {
	result, err := ls.repo.Update(id, input)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrLessonNotFound):
			return 0, apperror.New(http.StatusNotFound, "lesson not found", err)
		case errors.Is(err, models.ErrCourseNotFound):
			return 0, apperror.New(http.StatusNotFound, "course not found", err)
		default:
			return 0, apperror.New(http.StatusInternalServerError, "failed to update lesson", err)
		}
	}
	return result, nil
}
