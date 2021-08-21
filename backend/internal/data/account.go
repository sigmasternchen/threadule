package data

import (
	"gorm.io/gorm/clause"
	"threadule/backend/internal/data/models"
)

func (d *Data) GetAccountsByUser(user *models.User) ([]models.Account, error) {
	var accounts []models.Account
	err := d.db.
		Preload("Threads").
		Where("user_id = ?", user.ID).
		Where("access_token IS NOT NULL").
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

func (d *Data) GetAccountById(user *models.User, id string) (*models.Account, error) {
	var account models.Account
	err := d.db.
		Where("user_id = ?", user.ID).
		Where("id = ?", id).
		First(&account).
		Error
	if err != nil {
		return nil, err
	} else {
		return &account, nil
	}
}

func (d *Data) AddAccount(account *models.Account) error {
	return d.db.
		Omit(clause.Associations).
		Create(account).
		Error
}

func (d *Data) UpdateAccount(account *models.Account) error {
	return d.db.
		Omit(clause.Associations).
		Save(account).
		Error
}
