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
	case logic.ErrNotFound:
		StatusResponse(ctx, http.StatusNotFound, err.Error())
	case logic.ErrInvalidParameter:
		StatusResponse(ctx, http.StatusBadRequest, err.Error())
	case logic.ErrInsufficientPrivilege:
		StatusResponse(ctx, http.StatusForbidden, err.Error())
	case logic.ErrLoginFailed:
		StatusResponse(ctx, http.StatusForbidden, err.Error())
	case logic.ErrConflict:
		StatusResponse(ctx, http.StatusConflict, err.Error())
	case logic.ErrInvalidSession:
		StatusResponse(ctx, http.StatusUnauthorized, err.Error())
	case logic.ErrInternalError:
		StatusResponse(ctx, http.StatusInternalServerError, err.Error())
	default:
		StatusResponse(ctx, http.StatusInternalServerError, err.Error())
	}
}

func RedirectResponse(ctx *web.Context, url string) {
	ctx.Response.Header().Add("Location", url)
	ctx.Response.WriteHeader(http.StatusFound)
}
