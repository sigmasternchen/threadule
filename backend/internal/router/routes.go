package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
	. "threadule/backend/internal/presentation"
)

func Setup(ctx *app.Context) http.Handler {
	router := &router{Router: httprouter.New(), appCtx: ctx}

	router.POST("/authentication", Login)
	router.GET("/authentication", authenticated(GetAuthenticationData))

	router.GET("/account/", authenticated(GetAccounts))

	return router
}
