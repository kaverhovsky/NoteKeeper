package dto

type NoteDTO struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Author string   `json:"author"`
	Tags   []string `json:"tags"`
}
