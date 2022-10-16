package note

import "NoteKeeper/internal/user"

type Note struct {
	Title      string
	Body       string
	Maintainer *user.User
}
