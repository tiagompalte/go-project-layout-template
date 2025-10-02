package usecase

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
)

type FindAllNoteUseCase interface {
	Execute(ctx context.Context, input FindAllNoteInput) ([]entity.Note, error)
}

type FindAllNoteInput struct {
	Limit int64
}

type FindAllNoteUseCaseImpl struct {
	noteRepo protocols.NoteRepository
}

func NewFindAllNoteUseCaseImpl(noteRepo protocols.NoteRepository) FindAllNoteUseCase {
	return FindAllNoteUseCaseImpl{
		noteRepo: noteRepo,
	}
}

func (uc FindAllNoteUseCaseImpl) Execute(ctx context.Context, input FindAllNoteInput) ([]entity.Note, error) {
	notes, err := uc.noteRepo.FindAll(ctx, input.Limit)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return notes, nil
}
