package router

import (
	"net/http"
	"strings"
	. "threadule/backend/internal/presentation"
	"threadule/backend/internal/web"
)

const authPrefix = "Bearer "

func authenticated(next web.Handler) web.Handler {
	return func(ctx *web.Context) {
		authHeader := ctx.Request.Header.Get("Authentication")
		if !strings.HasPrefix(authHeader, authPrefix) {
			StatusResponse(ctx, http.StatusBadRequest, "Authentication header missing or malformed")
			return
		}
		authHeader = strings.TrimPrefix(authHeader, authPrefix)

		user, err := ctx.AppCtx.Logic.AuthenticateSession(authHeader)
		if err != nil {
			StatusResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}

		ctx.Session.User = user

		next(ctx)
	}
}
