package data

import "threadule/backend/internal/data/models"

func (d *Data) AddGroup(group *models.Group) error {
	return d.db.
		Create(group).
		Error
}
