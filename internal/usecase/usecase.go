package usecase

import (
	"NoteKeeper/internal/domain"
	"NoteKeeper/internal/domain/convert"
	"NoteKeeper/internal/domain/dto"
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

func (u *NoteUsecase) GetNotes(optsDTO dto.SearchOptionsDTO) ([]domain.Note, error) {
	opts, err := convert.DtoToOptions(optsDTO)
	if err != nil {
		u.logger.With(zap.NamedError("reason", err)).Error("failed to convert optionsDTO to options")
		return nil, err
	}
	return u.repo.Get(opts)
}

// TODO определить, где лучше логировать ошибки: в usecase или выше
func (u *NoteUsecase) Create(noteDTO dto.NoteDTO) (domain.Note, error) {
	note, err := convert.DtoToNote(noteDTO)
	if err != nil {
		u.logger.With(zap.NamedError("reason", err)).Error("failed to convert noteDTO to note")
		return domain.Note{}, err
	}

	if err := u.repo.Insert(note); err != nil {
		return domain.Note{}, err
	}

	return note, u.repo.Insert(note)
}

func (u *NoteUsecase) Update(noteDTO dto.NoteDTO) (domain.Note, error) {
	note, err := convert.DtoToNote(noteDTO)
	if err != nil {
		u.logger.With(zap.NamedError("reason", err)).Error("failed to convert noteDTO to note")
		return domain.Note{}, err
	}
	if err := u.repo.Update(note); err != nil {
		return domain.Note{}, err
	}

	return note, nil
}

func (u *NoteUsecase) Delete(uid uuid.UUID) error {
	return u.repo.Delete(uid)
}
