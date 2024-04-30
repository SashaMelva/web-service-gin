package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	config := config.New(configFile)
	log := logger.New(config.Logger, "../logFiles/")

	connectionDB := connection.New(config.DataBase, log)

	memstorage := memory.New(connectionDB.StorageDb, log)
	app := app.New(log, memstorage, config.HostClientApi)

	httpServer := http.NewServer(log, app, config.HttpServer)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			log.Error("failed to stop http server: " + err.Error())
		}
	}()

	log.Info("Services is running...")
	log.Debug("Debug mode enabled")

	if err := httpServer.Start(ctx); err != nil {
		log.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
