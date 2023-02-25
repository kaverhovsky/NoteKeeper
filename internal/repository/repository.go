package repository

import (
	"NoteKeeper/internal/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Init(dsn string) error
	Create(note domain.Note) error
	GetOne(uid uuid.UUID) error
	Get(opts domain.SearchOptions) ([]domain.Note, error)
	Update(note domain.Note) error
	Delete(uid uuid.UUID) error
}
