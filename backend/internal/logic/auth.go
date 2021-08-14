package logic

import (
	"errors"
	"threadule/backend/internal/data/models"
	"time"
)

const sessionDuration = 7 * 24 * time.Hour

func (l *Logic) scheduleTriggerAuth() {
	err := l.ctx.Data.CleanupSessions()
	if err != nil {
		l.ctx.Log.Errorf("couldn't clean up sessions: %v", err)
	}
}

func (l *Logic) AuthenticateSession(token string) (*models.User, error) {
	session, err := l.ctx.Data.GetSession(token)
	if err != nil {
		return nil, errors.New("invalid session")
	}

	session.ValidUntil = time.Now().Add(sessionDuration)
	err = l.ctx.Data.UpdateSession(session)
	if err != nil {
		l.ctx.Log.Errorf("couldn't extend session: %v", err)
	}

	return session.User, nil
}
