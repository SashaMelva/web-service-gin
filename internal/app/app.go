package app

import (
	storage "github.com/SashaMelva/web-service-gin/internal/memory/storage/postgre"
	"go.uber.org/zap"
)

type App struct {
	HostClientApi string
	storage       *storage.Storage
	Logger        *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, storage *storage.Storage, host *string) *App {
	return &App{
		HostClientApi: *host,
		storage:       storage,
		Logger:        logger,
	}
}
