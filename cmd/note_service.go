package main

import (
	v1 "NoteKeeper/internal/delivery/http/v1"
	"NoteKeeper/internal/repository"
	"NoteKeeper/internal/usecase"
	"NoteKeeper/pkg/common"
	"NoteKeeper/pkg/httpserver"
	"flag"
	"github.com/fasthttp/router"
	"go.uber.org/zap"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("c", ".env", "path to a config")
	flag.Parse()
	// instantiate startup logger before we read config
	startupLogger := common.NewLogger("development", "info")
	// read config according to provided path from the flag
	config, err := common.ReadConfig(*configPath, startupLogger)
	if err != nil {
		startupLogger.With(zap.NamedError("reason", err)).Fatal("failed to read the config")
	}
	// now we can create logger with desired parameters
	logger := common.NewLogger(config.Mode(), config.Level())

	// create and init repository
	repo := repository.NewPostgres(logger)
	if err := repo.Init(config.Postgres); err != nil {
		logger.With(zap.NamedError("reason", err)).Fatal("failed to init postgres database")
	}

	// create usecase instance
	usecase := usecase.NewNoteUsecase(logger, repo)

	// create new router
	router := router.New()
	server := httpserver.New(logger, router)

	// TODO добавить middlewares
	// TODO добавить statusAPI

	// create new note app
	apiV1 := v1.NewApiV1(logger, usecase)
	apiV1.AddRoutes(router.Group("/api/v1"))

	ln, err := net.Listen("tcp4", config.Listen)
	if err != nil {
		logger.With(zap.String("address", config.Listen),
			zap.NamedError("reason", err)).Fatal("failed to create tcp4 listener")
	}

	go server.Run(ln)

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-stopCh

	server.Shutdown()
	logger.Info("service stopped")
}
