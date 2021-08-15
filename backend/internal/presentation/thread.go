package presentation

import (
	"threadule/backend/internal/data/models"
	"threadule/backend/internal/web"
)

func AddThread(ctx *web.Context) {
	var thread models.Thread
	err := ctx.ReadJSON(&thread)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.AppCtx.Logic.AddThread(&thread, ctx.Session.User)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(&thread)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
}
