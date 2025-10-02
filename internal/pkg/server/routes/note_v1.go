package routes

import (
	"net/http"

	"github.com/tiagompalte/golang-clean-arch-template/application"
	"github.com/tiagompalte/golang-clean-arch-template/internal/pkg/server/handler"
	"github.com/tiagompalte/golang-clean-arch-template/pkg/server"
)

func CreateGroupLogV1(app application.App) server.GroupRoute {
	routes := []server.Route{
		{
			Path:    "/",
			Method:  http.MethodPost,
			Handler: handler.CreateNoteHandler(app.UseCase().CreateNoteUseCase),
		},
		{
			Path:    "/",
			Method:  http.MethodGet,
			Handler: handler.FindAllNoteHandler(app.UseCase().FindAllNoteUseCase),
		},
		{
			Path:    "/{id}",
			Method:  http.MethodGet,
			Handler: handler.FindByIDNoteHandler(app.UseCase().FindByIDNoteUseCase),
		},
	}

	return server.GroupRoute{
		Path:   "/notes",
		Routes: routes,
	}
}
