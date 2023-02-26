package dto

type SearchOptionsDTO struct {
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Tag    *string `json:"tags"`
}
