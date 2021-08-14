package logic

import (
	"net/url"
	"threadule/backend/internal/data/models"
)

func (l *Logic) GetAccounts(user *models.User) ([]models.Account, error) {
	return l.ctx.Data.GetAccountsByUser(user)
}

func (l *Logic) AddAccount(user *models.User) (string, *url.URL, error) {
	var account models.Account
	account.AccessToken = nil
	account.AccessTokenSecret = nil
	account.RequestSecret = nil
	account.UserID = user.ID

	err := l.ctx.Data.AddAccount(&account)
	if err != nil {
		l.ctx.Log.Errorf("couldn't create account in database: %v", err)
		return "", nil, ErrInternalError
	}

	return l.twitterLoginInit(&account)
}

func (l *Logic) AddAccountResolve(user *models.User, id string, pin string) (*models.Account, error) {
	account, err := l.ctx.Data.GetAccountById(user, id)
	if err != nil {
		l.ctx.Log.Errorf("couldn't get account for id: %v", err)
		return nil, ErrInvalidParameter
	}

	err = l.twitterLoginResolve(account, pin)
	if err != nil {
		return nil, err
	}

	return account, nil
}
