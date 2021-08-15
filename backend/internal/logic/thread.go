package logic

import (
	uuid "github.com/satori/go.uuid"
	"threadule/backend/internal/data/models"
)

func (l *Logic) AddThread(thread *models.Thread, user *models.User) error {
	if uuid.Nil == thread.AccountID {
		if thread.Account == nil {
			return ErrInvalidParameter
		}
		if uuid.Nil == thread.Account.ID {
			return ErrInvalidParameter
		}
		thread.AccountID = thread.Account.ID
	}

	_, err := l.ctx.Data.GetAccountById(user, thread.AccountID.String())
	if err != nil {
		return ErrNotFound
	}

	err = l.ctx.Data.AddThread(thread)
	return err
}
