package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"threadule/backend/internal/app"
	"threadule/backend/internal/presentation"
)

func Setup(ctx *app.Context) http.Handler {
	router := &router{Router: httprouter.New(), appCtx: ctx}

	router.GET("/", presentation.Test)

	return router
}
