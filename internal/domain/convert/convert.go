package convert

import (
	"NoteKeeper/internal/domain"
	"NoteKeeper/internal/domain/dto"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func DtoToNote(dto dto.NoteDTO) (domain.Note, error) {
	// generate new uid for note
	uid := uuid.New()
	//parse author uid

	author, err := uuid.Parse(dto.Author)
	if err != nil {
		return domain.Note{}, errors.Wrap(err, "failed to parse author uuid")
	}

	// parse tags array
	var tags []uuid.UUID
	for _, tagStr := range dto.Tags {
		tag, err := uuid.Parse(tagStr)
		if err != nil {
			return domain.Note{}, errors.Wrap(err, "failed to parse tag uuid")
		}

		tags = append(tags, tag)
	}

	return domain.Note{
		UID:    uid,
		Title:  dto.Title,
		Body:   dto.Body,
		Author: author,
		Tags:   tags,
	}, nil
}

func DtoToOptions(optsDto dto.SearchOptionsDTO) (domain.SearchOptions, error) {
	var opts domain.SearchOptions

	if optsDto.Title != nil {
		opts.Title = optsDto.Title
	}
	if optsDto.Author != nil {
		author, err := uuid.Parse(*optsDto.Author)
		if err != nil {
			return domain.SearchOptions{}, errors.Wrap(err, "failed to parse author uuid")
		}
		opts.Author = &author
	}

	if optsDto.Tag != nil {
		tag, err := uuid.Parse(*optsDto.Tag)
		if err != nil {
			return domain.SearchOptions{}, errors.Wrap(err, "failed to parse tag uuid")
		}
		opts.Tag = &tag
	}
	return opts, nil
}
