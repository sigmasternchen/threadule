package data

import (
	"threadule/backend/internal/data/models"
	"time"
)

func (d *Data) CleanupSessions() error {
	return d.db.
		Where("valid_until < ?", time.Now()).
		Delete(&models.Session{}).
		Error
}
