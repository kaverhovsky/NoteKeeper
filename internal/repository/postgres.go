package repository

import (
	"NoteKeeper/internal/domain"
	"github.com/google/uuid"
)

type Postgres struct {
}

func (p Postgres) Init(dsn string) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) Create(note domain.Note) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) GetOne(uid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) Get(opts domain.SearchOptions) ([]domain.Note, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) Update(note domain.Note) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) Delete(uid uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewPostgres() *Postgres {
	return &Postgres{}
}
