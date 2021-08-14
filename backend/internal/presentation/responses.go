package presentation

import (
	"net/http"
	"threadule/backend/internal/logic"
	"threadule/backend/internal/presentation/dto"
	"threadule/backend/internal/web"
	"time"
)

func StatusResponse(ctx *web.Context, status int, details string) {
	ctx.Response.WriteHeader(status)

	err := ctx.WriteJSON(&dto.Status{
		Status:  status,
		Message: http.StatusText(status),
		Details: details,
		Time:    time.Now(),
	})
	if err != nil {
		_, _ = ctx.Response.Write([]byte("something went very wrong"))
	}
}

func ErrorResponse(ctx *web.Context, err error) {
	switch err {
	case logic.ErrLoginFailed:
		StatusResponse(ctx, http.StatusForbidden, err.Error())
	case logic.ErrInvalidSession:
		StatusResponse(ctx, http.StatusUnauthorized, err.Error())
	case logic.ErrInternalError:
		StatusResponse(ctx, http.StatusInternalServerError, err.Error())
	default:
		StatusResponse(ctx, http.StatusInternalServerError, err.Error())
	}
}
