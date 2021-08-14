package app

import "threadule/backend/internal/data/models"

type Data interface {
	CountUsers() (int64, error)
	AddUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
	AddUserToGroup(user *models.User, group *models.Group) error
	DeleteUserFromGroup(user *models.User, group *models.Group) error

	AddGroup(group *models.Group) error

	GetSession(id string) (*models.Session, error)
	AddSession(session *models.Session) error
	UpdateSession(session *models.Session) error
	CleanupSessions() error

	GetAccountsByUser(user *models.User) ([]models.Account, error)
	GetAccountById(user *models.User, id string) (*models.Account, error)
	AddAccount(account *models.Account) error
	UpdateAccount(account *models.Account) error

	GetScheduledThreads() ([]models.Thread, error)
	GetTweetsForThread(thread *models.Thread) ([]models.Tweet, error)
	UpdateThread(thread *models.Thread) error

	UpdateTweet(tweet *models.Tweet) error
}
