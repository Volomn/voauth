package note

import "net/http"

type AddNoteRequestModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (a *AddNoteRequestModel) Bind(r *http.Request) error {
	return nil
}
