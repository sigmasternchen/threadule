package data

import (
	"gorm.io/gorm/clause"
	"threadule/backend/internal/data/models"
)

func (d *Data) UpdateTweet(tweet *models.Tweet) error {
	return d.db.
		Omit(clause.Associations).
		Save(tweet).
		Error
}
