package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type NoteResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   any       `json:"message"`
}

// @Summary Find All Notes
// @Description Find all notes
// @Tags Note
// @Produce json
// @Param limit query int false "Limit" default(100)
// @Success 200 {object} []NoteResponse "Notes list"
// @Router /api/v1/notes [get]
func FindAllNoteHandler(findAllNoteUseCase usecase.FindAllNoteUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		limit := int64(100) // default limit

		if v := r.URL.Query().Get("limit"); v != "" {
			if i, err := strconv.ParseInt(v, 10, 64); err == nil {
				limit = i
			}
		}

		input := usecase.FindAllNoteInput{
			Limit: limit,
		}

		notes, err := findAllNoteUseCase.Execute(ctx, input)
		if err != nil {
			return errors.Wrap(err)
		}

		resp := make([]NoteResponse, len(notes))
		for i := range notes {
			resp[i] = NoteResponse{
				ID:        notes[i].ID,
				CreatedAt: notes[i].CreatedAt,
				Message:   notes[i].Message,
			}
		}

		err = server.RespondJSON(w, http.StatusOK, resp)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
