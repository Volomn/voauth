package note

import (
	"net/http"

	"github.com/Volomn/voauth/backend/queryservice/note"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

type NoteListItemResponse struct {
	UUID       uuid.UUID `json:"uuid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IsPublic   bool      `json:"isPublic"`
	IsFavorite bool      `json:"isFavorite"`
	IsArchived bool      `json:"isArchived"`
}

func (rd *NoteListItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewNoteListResponse(items []note.NoteListItem) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		list = append(list, &NoteListItemResponse{
			UUID:       item.UUID,
			Title:      item.Title,
			Content:    item.Content,
			IsPublic:   item.IsPublic,
			IsFavorite: item.IsFavorite,
			IsArchived: item.IsArchived,
		})
	}
	return list
}

type NoteResponse struct {
	UUID       uuid.UUID `json:"uuid"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	IsPublic   bool      `json:"isPublic"`
	IsFavorite bool      `json:"isFavorite"`
	IsArchived bool      `json:"isArchived"`
}

func (rd *NoteResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewNoteResponse(note note.Note) render.Renderer {
	return &NoteResponse{
		UUID:       note.UUID,
		Title:      note.Title,
		Content:    note.Content,
		IsPublic:   note.IsPublic,
		IsFavorite: note.IsFavorite,
		IsArchived: note.IsArchived,
	}
}
