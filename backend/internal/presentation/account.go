package presentation

import (
	"threadule/backend/internal/presentation/dto"
	"threadule/backend/internal/web"
)

func GetAccounts(ctx *web.Context) {
	accounts, err := ctx.AppCtx.Logic.GetAccounts(ctx.Session.User)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(accounts)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
}

func AddAccount(ctx *web.Context) {
	id, url, err := ctx.AppCtx.Logic.AddAccount(ctx.Session.User)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(&dto.AddAccountResponse{
		ID:  id,
		URL: url.String(),
	})
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
}

func AddAccountResolve(ctx *web.Context) {
	id := ctx.Params.ByName("id")

	var param dto.AddAccountResolveParam
	err := ctx.ReadJSON(&param)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	account, err := ctx.AppCtx.Logic.AddAccountResolve(ctx.Session.User, id, param.Pin)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(account)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
}
