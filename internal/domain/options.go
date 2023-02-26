package domain

import "github.com/google/uuid"

// SearchOptions is an entity for searching records in storage
type SearchOptions struct {
	Title  *string
	Author *uuid.UUID
	// TODO пока что будем совершать поиск по одному тегу
	// TODO добавить поиск по нескольким (&& и ||)
	Tag *uuid.UUID
}
