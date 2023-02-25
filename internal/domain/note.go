package domain

import "github.com/google/uuid"

// Note is Main entity for Note Service
// contains Title, Body, Author (UserService), Tags(TagService)
// Now I will not separate domain and storage entities, it wil be the same one
type Note struct {
	tableName struct{}    `pg:"notes"`
	UID       uuid.UUID   `json:"uid" pg:"uid, type:uuid, pk"`
	Title     string      `json:"title" pg:"title, use_zero"`
	Body      string      `json:"body" pg:"body, use_zero"`
	Author    uuid.UUID   `json:"author" pg:"author, type:uuid, use_zero"`
	Tags      []uuid.UUID `json:"tags" pg:"tags, type:array"`
}
