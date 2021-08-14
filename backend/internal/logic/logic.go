package logic

import (
	"threadule/backend/internal/app"
)

type Logic struct {
	ctx *app.Context
}

var _ app.Logic = &Logic{}

func Setup(ctx *app.Context) (app.Logic, error) {
	logic := &Logic{
		ctx: ctx,
	}

	logic.startScheduler()

	return logic, nil
}
