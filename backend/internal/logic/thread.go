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

	thread.Status = models.ThreadScheduled
	for i := range thread.Tweets {
		thread.Tweets[i].Status = models.TweetScheduled
	}

	err = l.ctx.Data.AddThread(thread)
	return err
}

func (l *Logic) UpdateThread(thread *models.Thread, user *models.User) error {
	oldThread, err := l.ctx.Data.GetThread(thread.ID, user)
	if err != nil {
		return ErrNotFound
	}
	thread.AccountID = oldThread.AccountID

	thread.Status = models.ThreadScheduled
	for i := range thread.Tweets {
		thread.Tweets[i].Status = models.TweetScheduled
	}

	err = l.ctx.Data.UpdateThread(thread)
	return err
}

func (l *Logic) DeleteThread(id uuid.UUID) error {
	return l.ctx.Data.DeleteThread(id)
}

func (l *Logic) GetThreads(user *models.User) ([]models.Thread, error) {
	return l.ctx.Data.GetThreads(user)
}
