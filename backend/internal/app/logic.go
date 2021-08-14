package app

import "threadule/backend/internal/data/models"

type Logic interface {
	AuthenticateSession(token string) (*models.User, error)
}
