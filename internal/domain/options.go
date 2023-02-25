package domain

import "github.com/google/uuid"

// SearchOptions is an entity for searching records in storage
type SearchOptions struct {
	Title  *string
	Author *uuid.UUID
	Tags   []uuid.UUID
}
