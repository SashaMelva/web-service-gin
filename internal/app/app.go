package app

import (
	"github.com/SashaMelva/web-service-gin/internal/entity"
	storage "github.com/SashaMelva/web-service-gin/internal/memory/storage/postgre"
	"go.uber.org/zap"
)

type App struct {
	storage *storage.Storage
	log     *zap.SugaredLogger
	period  entity.Period
}

func New(logger *zap.SugaredLogger, storage *storage.Storage) *App {
	return &App{
		storage: storage,
		log:     logger,
		period: map[string]string{
			"week":   "week",
			"mounth": "mounth",
			"today":  "mounth",
		},
	}
}
