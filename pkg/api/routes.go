package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (api *Api) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/todos", api.handleTodoList)
	r.Get("/todos/{id}", api.handleTodoGet)
	r.Post("/todos", api.handleTodoCreate)
	r.Put("/todos/{id}", api.handleTodoUpdate)
	r.Delete("/todos/{id}", api.handleTodoDelete)

	return r
}

func (api *Api) getID(r *http.Request) (int, error) {
	idString := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		return 0, err
	}
	return id, nil
}
