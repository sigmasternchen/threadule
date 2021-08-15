package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
	. "threadule/backend/internal/presentation"
)

func Setup(ctx *app.Context) http.Handler {
	router := &router{Router: httprouter.New(), appCtx: ctx}

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx.AccessLog.Printf("%s %s; not found", request.Method, request.URL.String())
		writer.WriteHeader(http.StatusNotFound)
	})

	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx.AccessLog.Printf("%s %s; method not allowed", request.Method, request.URL.String())
		writer.WriteHeader(http.StatusMethodNotAllowed)
	})

	router.POST("/authentication", Login)
	router.GET("/authentication", authenticated(GetAuthenticationData))

	router.GET("/account/", authenticated(GetAccounts))
	router.POST("/account/", authenticated(AddAccount))
	router.POST("/account/:id", authenticated(AddAccountResolve))

	router.POST("/thread/", authenticated(AddThread))

	return router
}
