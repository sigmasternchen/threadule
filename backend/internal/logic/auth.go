package logic

import (
	"errors"
	"threadule/backend/internal/data/models"
	"time"
)

const sessionDuration = 7 * 24 * time.Hour

var (
	ErrLoginFailed    = errors.New("login failed")
	ErrInvalidSession = errors.New("invalid session")
	ErrInternalError  = errors.New("something went wrong")
)

func (l *Logic) scheduleTriggerAuth() {
	err := l.ctx.Data.CleanupSessions()
	if err != nil {
		l.ctx.Log.Errorf("couldn't clean up sessions: %v", err)
	}
}

func (l *Logic) AuthenticateSession(token string) (*models.User, error) {
	session, err := l.ctx.Data.GetSession(token)
	if err != nil {
		return nil, ErrInvalidSession
	}

	session.ValidUntil = time.Now().Add(sessionDuration)
	err = l.ctx.Data.UpdateSession(session)
	if err != nil {
		l.ctx.Log.Errorf("couldn't extend session: %v", err)
	}

	return session.User, nil
}

func (l *Logic) Login(username, password string) (string, error) {
	user, err := l.ctx.Data.GetUserByUsername(username)

	// the following few lines should prevent timing attacks
	hash := "something"
	if err == nil {
		hash = user.Password
	} else {
		password = "something else"
	}

	if l.checkPassword(hash, password) && user != nil {
		session := &models.Session{
			UserID:     user.ID,
			ValidUntil: time.Now().Add(sessionDuration),
		}
		err = l.ctx.Data.AddSession(session)
		if err != nil {
			l.ctx.Log.Errorf("couldn't create session for '%v': %v", user.Username, err)
			return "", ErrInternalError
		}

		return session.ID.String(), nil
	} else {
		return "", ErrLoginFailed
	}
}
