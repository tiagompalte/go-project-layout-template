package handler

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

// @Summary Find Note by ID
// @Description Find Note by ID
// @Tags Note
// @Produce json
// @Param id path string true "Note ID"
// @Success 200 {object} NoteResponse "Note found"
// @Router /api/v1/notes/{id} [get]
func FindByIDNoteHandler(findByIDNoteUseCase usecase.FindByIDNoteUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		id := r.PathValue("id")

		note, err := findByIDNoteUseCase.Execute(ctx, id)
		if err != nil {
			return errors.Wrap(err)
		}

		var response NoteResponse
		response.ID = note.ID
		response.CreatedAt = note.CreatedAt
		response.Message = note.Message

		err = server.RespondJSON(w, http.StatusOK, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
