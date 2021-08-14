package presentation

import (
	"net/http"
	"threadule/backend/internal/web"
)

func GenericStatusResponse(ctx *web.Context, status int) {
	ctx.Response.WriteHeader(status)
	_, _ = ctx.Response.Write([]byte(http.StatusText(status)))
}
