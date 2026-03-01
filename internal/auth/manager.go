package auth

import "github.com/KozhabergenovNurzhan/GoProj1/internal/models"

type TokenManager interface {
	NewAccessToken(user models.User) (string, int64, error)
}
