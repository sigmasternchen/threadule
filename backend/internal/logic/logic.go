package logic

import (
	"threadule/backend/internal/app"
)

type Logic struct {
	ctx *app.Context
}

var _ app.Logic = &Logic{}
