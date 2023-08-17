package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mrtj458/bolttodo/pkg/data"
)

func (api *Api) handleTodoList(w http.ResponseWriter, r *http.Request) {
	todos, err := api.TodoService.GetAll()
	if err != nil {
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	api.writeJSON(w, r, http.StatusOK, todos)
}

func (api *Api) handleTodoGet(w http.ResponseWriter, r *http.Request) {
	id, err := api.getID(r)
	if err != nil {
		api.notFoundResponse(w, r)
		return
	}

	t, err := api.TodoService.GetByID(id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			api.notFoundResponse(w, r)
			return
		}
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	api.writeJSON(w, r, http.StatusOK, t)
}

func (api *Api) handleTodoCreate(w http.ResponseWriter, r *http.Request) {
	var t data.Todo
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		api.Logger.Error(err.Error())
		api.internalErrorResponse(w, r)
		return
	}

	if len(t.Value) == 0 {
		api.writeJSON(w, r, http.StatusUnprocessableEntity, map[string]string{"error": "value must be provided"})
		return
	}

	err = api.TodoService.Insert(&t)
	if err != nil {
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	api.writeJSON(w, r, http.StatusCreated, t)
}

func (api *Api) handleTodoUpdate(w http.ResponseWriter, r *http.Request) {
	id, err := api.getID(r)
	if err != nil {
		api.notFoundResponse(w, r)
		return
	}

	var t data.Todo
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	if len(t.Value) == 0 {
		api.writeJSON(w, r, http.StatusUnprocessableEntity, map[string]string{"error": "value must be provided"})
		return
	}

	err = api.TodoService.Update(id, &t)
	if err != nil {
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	api.writeJSON(w, r, http.StatusCreated, t)
}

func (api *Api) handleTodoDelete(w http.ResponseWriter, r *http.Request) {
	id, err := api.getID(r)
	if err != nil {
		api.notFoundResponse(w, r)
		return
	}

	err = api.TodoService.Delete(id)
	if err != nil {
		if errors.Is(err, data.ErrNotFound) {
			api.notFoundResponse(w, r)
			return
		}
		api.internalErrorResponse(w, r, "err", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
