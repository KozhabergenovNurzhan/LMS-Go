package repository

import "github.com/KozhabergenovNurzhan/GoProj1/internal/config"

// fake DB (later implement sqlx.DB)
type DB struct {
	connectionPath string
}

func NewPostgresDB(cfg *config.Config) (*DB, error) {
	db := &DB{}
	return db, nil
}
