package data

import (
	"gorm.io/gorm/clause"
	"threadule/backend/internal/data/models"
)

func (d *Data) CountUsers() (int64, error) {
	var c int64
	err := d.db.
		Model(&models.User{}).
		Count(&c).
		Error
	return c, err
}

func (d *Data) CreateUser(user *models.User) error {
	return d.db.
		Omit(clause.Associations).
		Create(user).
		Error
}

func (d *Data) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := d.db.
		Where("username = ?", username).
		First(&user).
		Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
}
