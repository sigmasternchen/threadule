package presentation

import (
	"threadule/backend/internal/data/models"
	"threadule/backend/internal/web"
)

func UpdateSelf(ctx *web.Context) {
	var user models.User
	err := ctx.ReadJSON(&user)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	user.ID = ctx.Session.User.ID

	err = ctx.AppCtx.Logic.UpdateUser(&user, ctx.Session.User)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}

	err = ctx.WriteJSON(&user)
	if err != nil {
		ErrorResponse(ctx, err)
		return
	}
}

func GetSelf(ctx *web.Context) {
	err := ctx.WriteJSON(ctx.Session.User)
	if err != nil {
		ErrorResponse(ctx, err)
	}
}
