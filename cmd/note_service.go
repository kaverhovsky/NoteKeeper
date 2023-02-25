package main

import (
	"NoteKeeper/internal/domain"
	"NoteKeeper/internal/repository"
	"NoteKeeper/internal/usecase"
	"NoteKeeper/pkg/common"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
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

	note := domain.Note{
		UID:    uuid.New(),
		Title:  "my_note",
		Body:   "that's a body",
		Author: uuid.Nil,
		Tags:   []uuid.UUID{uuid.New()},
	}

	if err := usecase.Create(note); err != nil {
		logger.With(zap.NamedError("reason", err)).Error("failed to create note")
		return
	}

	if nt, err := usecase.GetNote(note.UID); err != nil {
		logger.With(zap.NamedError("reason", err)).Error("failed to create note")
	} else {
		logger.With(zap.String("note_instance", fmt.Sprintf("%v", nt))).Info("that's the note")
	}

	// TODO проверить апдейт и удаление

}
