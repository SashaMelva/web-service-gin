package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/SashaMelva/web-service-gin/internal/app"
	"github.com/SashaMelva/web-service-gin/internal/config"
	"github.com/SashaMelva/web-service-gin/internal/logger"
	"github.com/SashaMelva/web-service-gin/internal/memory/connection"
	storage "github.com/SashaMelva/web-service-gin/internal/memory/storage/postgre"
	"github.com/SashaMelva/web-service-gin/internal/server/http"
)

func main() {
	configFile := "../configs/"
	config := config.New(configFile)
	log := logger.New(config.Logger, "../logs/")

	connectionDB := connection.New(config.DataBase, log)
	// err := migrator.RunMigrationsPg(connectionDB, "migrations")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	memstorage := storage.New(connectionDB.StorageDb, log)
	app := app.New(log, memstorage)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		httpServer.Stop(ctx)
	}()

	log.Info("Services is running...")
	log.Debug("Debug mode enabled")

	httpServer.Start(ctx)
}
