package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/KozhabergenovNurzhan/GoProj1/internal/config"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/handler"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/pkg/logger"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/repository"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/server"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		return
	}

	courseRepo := repository.NewPsgCourseRepo(db)
	courseService := service.NewCourseService(courseRepo)

	h := handler.NewHandler(courseService)

	router, err := h.InitRoutes()

	srv := server.New(cfg.Port, router)
	err = srv.Run()
	if err != nil {
		slog.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
