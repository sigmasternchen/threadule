package data

import "threadule/backend/internal/data/models"

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
		Create(user).
		Error
}
