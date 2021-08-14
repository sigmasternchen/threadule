package router

import (
	"net/http"
	"strings"
	"threadule/backend/internal/presentation"
	"threadule/backend/internal/web"
)

const authPrefix = "Bearer "

func authenticated(next web.Handler) web.Handler {
	return func(ctx *web.Context) {
		authHeader := ctx.Request.Header.Get("Authentication")
		if !strings.HasPrefix(authHeader, authPrefix) {
			presentation.GenericStatusResponse(ctx, http.StatusUnauthorized)
			return
		}
		authHeader = strings.TrimPrefix(authHeader, authPrefix)

		user, err := ctx.AppCtx.Logic.AuthenticateSession(authHeader)
		if err != nil {
			presentation.GenericStatusResponse(ctx, http.StatusUnauthorized)
			return
		}

		ctx.Session.User = user

		next(ctx)
	}
}
