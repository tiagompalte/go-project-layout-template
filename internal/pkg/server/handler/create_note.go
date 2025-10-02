package handler

import (
	"encoding/json"
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/usecase"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

type CreateNoteRequest struct {
	Message any `json:"message"`
}

func (r *CreateNoteRequest) toInput() usecase.CreateNoteInput {
	return usecase.CreateNoteInput{
		Message: r.Message,
	}
}

// @Summary Create Note
// @Description Create new Note
// @Tags Note
// @Accept json
// @Produce json
// @Param new_note body CreateNoteRequest true "New Note"
// @Success 201 {object} NoteResponse "Create Note success"
// @Router /api/v1/notes [post]
func CreateNoteHandler(createNoteUseCase usecase.CreateNoteUseCase) server.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()

		var request CreateNoteRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return errors.Wrap(err)
		}

		output, err := createNoteUseCase.Execute(ctx, request.toInput())
		if err != nil {
			return errors.Wrap(err)
		}

		response := NoteResponse{
			ID:        output.ID,
			CreatedAt: output.CreatedAt,
			Message:   output.Message,
		}

		err = server.RespondJSON(w, http.StatusCreated, response)
		if err != nil {
			return errors.Wrap(err)
		}

		return nil
	}
}
