package usecase

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type CreateNoteInput struct {
	Message any
}

type CreateNoteUseCase interface {
	Execute(ctx context.Context, input CreateNoteInput) (entity.Note, error)
}

type CreateNoteUseCaseImpl struct {
	noteRepo protocols.NoteRepository
}

func NewCreateNoteUseCaseImpl(noteRepo protocols.NoteRepository) CreateNoteUseCase {
	return CreateNoteUseCaseImpl{
		noteRepo: noteRepo,
	}
}

func (uc CreateNoteUseCaseImpl) Execute(ctx context.Context, input CreateNoteInput) (entity.Note, error) {
	var note entity.Note
	note.CreatedAt = time.Now()
	note.Message = input.Message

	id, err := uc.noteRepo.Insert(ctx, note)
	if err != nil {
		return entity.Note{}, errors.Wrap(err)
	}

	note, err = uc.noteRepo.FindByID(ctx, id)
	if err != nil {
		return entity.Note{}, errors.Wrap(err)
	}

	return note, nil
}
