package data

import (
	"fmt"
	"gorm.io/gorm"
	"threadule/backend/internal/app"
)

type Data struct {
	ctx *app.Context
	db  *gorm.DB
}

var _ app.Data = &Data{}

func Setup(ctx *app.Context) (app.Data, error) {
	db, err := connect(ctx.Config.Database.DSN)
	if err != nil {
		ctx.Log.Errorf("failed to connect to database: %v", err)
		return nil, err
	}

	err = migrate(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to migrate models: %w", err)
	}

	return &Data{
		ctx: ctx,
		db:  db,
	}, nil
}
