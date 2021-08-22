package presentation

import (
	"net/http"
	"threadule/backend/internal/presentation/dto"
	"threadule/backend/internal/web"
)

func Login(ctx *web.Context) {
	var param dto.LoginParams
	err := ctx.ReadJSON(&param)
	if err != nil {
		StatusResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	sessionToken, err := ctx.AppCtx.Logic.Login(param.Username, param.Password)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(&dto.LoginResponse{
		Token: sessionToken,
	})
	if err != nil {
		ErrorResponse(ctx, err)
	}
}
