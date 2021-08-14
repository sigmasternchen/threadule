package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
	"threadule/backend/internal/data/models"
)

type SessionInfo struct {
	User *models.User
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	AppCtx   *app.Context
	Session  SessionInfo
}
