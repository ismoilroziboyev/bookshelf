package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/repository"
	"github.com/ismoilroziboyev/bookshelf/internal/services"
	"github.com/ismoilroziboyev/bookshelf/internal/transport/rest"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Info("Bookshelf project is starting...")

	cfg := config.Load()

	logrus.Info("configs loaded successfully....")

	repo := repository.New(cfg)

	logrus.Info("repository initialized successfully...")

	service := services.New(cfg, repo)

	logrus.Info("services initialized successfully...")

	restServer := rest.New(cfg, service)

	logrus.Info("rest server initialized successfully...")

	restErr := make(chan error, 1)

	go func() {
		logrus.Info("rest server runned successfully....")
		if err := restServer.Run(); err != nil {
			restErr <- err
		}
	}()

	q := make(chan os.Signal, 1)

	signal.Notify(q, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-q:
		logrus.Info("shutdown signal received from machine....")
	case err := <-restErr:
		logrus.Errorf("error occurred in rest server: %s", err.Error())
	}

	/// gracefully shutdown
	logrus.Info("Bookshelf project shutting down....")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := restServer.Shutdown(ctx); err != nil {
		logrus.Errorf("error occurred while shutting down rest server....")
	}

	if err := repo.Close(ctx); err != nil {
		logrus.Errorf("error occurred while closing repository connections...")
	}

	logrus.Info("Bookshelf project shutted down...")

}
