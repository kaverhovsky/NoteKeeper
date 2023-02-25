package usecase

import (
	"NoteKeeper/internal/domain"
	"NoteKeeper/internal/repository"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type NoteUsecase struct {
	logger *zap.Logger
	repo   repository.Repository
}

func NewNoteUsecase(logger *zap.Logger, repo repository.Repository) *NoteUsecase {
	return &NoteUsecase{
		logger: logger,
		repo:   repo,
	}
}

func (u *NoteUsecase) GetNote(uid uuid.UUID) (domain.Note, error) {
	return u.repo.GetOne(uid)
}

func (u *NoteUsecase) GetNotes(opts domain.SearchOptions) ([]domain.Note, error) {
	return u.repo.Get(opts)
}

func (u *NoteUsecase) Create(note domain.Note) error {
	return u.repo.Insert(note)
}

func (u *NoteUsecase) Update(note domain.Note) error {
	return u.repo.Update(note)
}

func (u *NoteUsecase) Delete(uid uuid.UUID) error {
	return u.repo.Delete(uid)
}
