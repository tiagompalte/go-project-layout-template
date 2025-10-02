package mongo

import (
	"context"
	"time"

	"github.com/tiagompalte/golang-clean-arch-template/internal/app/entity"
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/errors"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type noteDocument struct {
	ID        bson.ObjectID `bson:"_id,omitempty"`
	CreatedAt time.Time     `bson:"created_at"`
	Message   any           `bson:"message"`
}

type NoteRepository struct {
	conn       repository.ConnectorMongo
	collection string
}

func NewNoteRepository(conn repository.ConnectorMongo) protocols.NoteRepository {
	return NoteRepository{
		conn:       conn,
		collection: "notes",
	}
}

func (r NoteRepository) Insert(ctx context.Context, note entity.Note) (string, error) {
	noteDoc := noteDocument{
		CreatedAt: note.CreatedAt,
		Message:   note.Message,
	}

	res, err := r.conn.InsertOne(ctx, r.collection, noteDoc)
	if err != nil {
		return "", errors.Wrap(err)
	}

	return res.InsertedID.Hex(), nil
}

func (r NoteRepository) FindAll(ctx context.Context, limit int64) ([]entity.Note, error) {
	result, err := r.conn.Find(ctx, r.collection, bson.D{}, options.Find().SetLimit(limit))
	if err != nil {
		return nil, errors.Wrap(err)
	}

	var notes []entity.Note
	for result.Next(ctx) {
		var doc noteDocument
		if err := result.Decode(&doc); err != nil {
			return nil, errors.Wrap(err)
		}

		note := entity.Note{
			ID:        doc.ID.Hex(),
			CreatedAt: doc.CreatedAt,
			Message:   doc.Message,
		}

		notes = append(notes, note)
	}

	return notes, nil
}

func (r NoteRepository) FindByID(ctx context.Context, id string) (entity.Note, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return entity.Note{}, errors.Wrap(err)
	}

	var doc noteDocument
	err = r.conn.FindOne(ctx, r.collection, bson.M{"_id": objectID}, &doc)
	if err != nil {
		return entity.Note{}, errors.Wrap(err)
	}

	return entity.Note{
		ID:        doc.ID.Hex(),
		CreatedAt: doc.CreatedAt,
		Message:   doc.Message,
	}, nil
}
