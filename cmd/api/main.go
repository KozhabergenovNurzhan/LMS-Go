package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/KozhabergenovNurzhan/GoProj1/internal/config"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/pkg/logger"
	"github.com/KozhabergenovNurzhan/GoProj1/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err.Error())
	}

	slogger := logger.New(cfg.LogLevel)
	slog.SetDefault(slogger)

	r := gin.New()

	srv := server.New(cfg.Port, r)
	err = srv.Run()
	if err != nil {
		slog.Error("failed to start server", "error", err.Error())
		os.Exit(1)
	}
}
