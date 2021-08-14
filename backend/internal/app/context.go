package app

import "threadule/backend/internal/config"

type Context struct {
	Config *config.Config
	Logic  Logic
	Data   Data
}
