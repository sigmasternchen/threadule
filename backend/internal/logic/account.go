package logic

import "threadule/backend/internal/data/models"

func (l *Logic) GetAccounts(user *models.User) ([]models.Account, error) {
	return l.ctx.Data.GetAccountsByUser(user)
}
