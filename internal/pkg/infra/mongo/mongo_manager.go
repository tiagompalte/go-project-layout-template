package mongo

import (
	"github.com/tiagompalte/golang-clean-arch-template/internal/app/protocols"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/infra/data"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/repository"
)

type manager struct {
	note protocols.NoteRepository
}

func NewRepositoryManager(conn repository.ConnectorMongo) data.Manager[data.MongoManager] {
	return manager{
		note: NewNoteRepository(conn),
	}
}

func (m manager) Repository() data.MongoManager {
	return data.MongoManager{
		Note: func() protocols.NoteRepository { return m.note },
	}
}
