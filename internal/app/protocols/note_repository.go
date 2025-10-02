package protocols

import (
	"context"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
)

type NoteRepository interface {
	Insert(ctx context.Context, note entity.Note) (string, error)
	FindAll(ctx context.Context, limit int64) ([]entity.Note, error)
	FindByID(ctx context.Context, id string) (entity.Note, error)
}
