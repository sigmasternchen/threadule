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

func (d *Data) GetSession(id string) (*models.Session, error) {
	var session models.Session
	err := d.db.
		Where("valid_until > ?", time.Now()).
		Where("id = ?", id).
		First(&session).
		Error
	if err != nil {
		return nil, err
	} else {
		return &session, nil
	}
}

func (d *Data) UpdateSession(session *models.Session) error {
	return d.db.
		Save(session).
		Error
}

func (d *Data) AddSession(session *models.Session) error {
	return d.db.
		Create(session).
		Error
}
