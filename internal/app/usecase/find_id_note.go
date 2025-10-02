package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindByIDNoteUseCase interface {
	Execute(ctx context.Context, id string) (entity.Note, error)
}

type FindByIDNoteUseCaseImpl struct {
	noteRepo protocols.NoteRepository
}

func NewFindByIDNoteUseCaseImpl(noteRepo protocols.NoteRepository) FindByIDNoteUseCase {
	return FindByIDNoteUseCaseImpl{noteRepo: noteRepo}
}

func (uc FindByIDNoteUseCaseImpl) Execute(ctx context.Context, id string) (entity.Note, error) {
	note, err := uc.noteRepo.FindByID(ctx, id)
	if err != nil {
		return entity.Note{}, errors.Wrap(err)
	}
	return note, nil
}
