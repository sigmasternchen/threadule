package data

import "threadule/backend/internal/data/models"

func (d *Data) GetAccountsByUser(user *models.User) ([]models.Account, error) {
	var accounts []models.Account
	err := d.db.
		Where("user_id = ?", user.ID).
		Find(&accounts).
		Error
	if err != nil {
		return nil, err
	} else {
		for i := range accounts {
			accounts[i].AccessToken = nil
			accounts[i].AccessTokenSecret = nil
		}

		return accounts, nil
	}
}
