package app

import (
	storage "github.com/SashaMelva/web-service-gin/internal/memory/storage/postgre"
	"go.uber.org/zap"
)

type App struct {
	storage *storage.Storage
	log     *zap.SugaredLogger
}

func New(logger *zap.SugaredLogger, storage *storage.Storage) *App {
	return &App{
		storage: storage,
		log:     logger,
	}
}
