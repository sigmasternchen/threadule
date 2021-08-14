package web

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	AppCtx   *app.Context
}
