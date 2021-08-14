package app

import (
	"github.com/google/logger"
	"log"
	"threadule/backend/internal/config"
)

type Context struct {
	Config    *config.Config
	Log       *logger.Logger
	AccessLog *log.Logger
	Logic     Logic
	Data      Data
}
