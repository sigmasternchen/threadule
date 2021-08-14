package presentation

import "threadule/backend/internal/web"

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
