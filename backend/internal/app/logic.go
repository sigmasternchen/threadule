package app

import (
	uuid "github.com/satori/go.uuid"
	"net/url"
	"threadule/backend/internal/data/models"
)

type Logic interface {
	UpdateUser(userToUpdate *models.User, currentUser *models.User) error

	AuthenticateSession(token string) (*models.User, error)
	Login(username, password string) (string, error)

	GetAccounts(user *models.User) ([]models.Account, error)
	AddAccount(user *models.User) (string, *url.URL, error)
	AddAccountResolve(user *models.User, id string, pin string) (*models.Account, error)

	AddThread(thread *models.Thread, user *models.User) error
	UpdateThread(thread *models.Thread, user *models.User) error
	DeleteThread(id uuid.UUID, user *models.User) error
	GetThreads(user *models.User) ([]models.Thread, error)
}
