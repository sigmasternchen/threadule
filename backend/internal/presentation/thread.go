package presentation

import (
	uuid "github.com/satori/go.uuid"
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

func UpdateThread(ctx *web.Context) {
	var thread models.Thread
	err := ctx.ReadJSON(&thread)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	thread.ID, err = uuid.FromString(ctx.Params.ByName("id"))
	if err != nil {
		ErrorResponse(ctx, err)
	}

	err = ctx.AppCtx.Logic.UpdateThread(&thread, ctx.Session.User)
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
