package data

import (
	uuid "github.com/satori/go.uuid"
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

func (d *Data) AddUser(user *models.User) error {
	return d.db.
		Omit(clause.Associations).
		Create(user).
		Error
}

func (d *Data) UpdateUser(user *models.User) error {
	return d.db.
		Omit(clause.Associations).
		Save(user).
		Error
}

func (d *Data) GetUser(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := d.db.
		Preload("Groups").
		First(&user, id).
		Error
	if err != nil {
		return nil, err
	} else {
		return &user, nil
	}
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

func (d *Data) AddUserToGroup(user *models.User, group *models.Group) error {
	return d.db.
		Model(user).
		Omit("Groups.*").
		Association("Groups").
		Append([]models.Group{*group})
}

func (d *Data) DeleteUserFromGroup(user *models.User, group *models.Group) error {
	return d.db.
		Model(user).
		Association("Groups").
		Delete(group)
}
