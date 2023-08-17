package api

import (
	"encoding/json"
	"net/http"
)

func (api *Api) writeJSON(w http.ResponseWriter, r *http.Request, statusCode int, data any) {
	jsn, err := json.Marshal(data)
	if err != nil {
		api.Logger.Error(err.Error(), "path", r.URL.Path, "method", r.Method)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsn)
	if err != nil {
		api.internalErrorResponse(w, r, "err", err, "path", r.URL.Path, "method", r.Method)
		return
	}
}

func (api *Api) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	api.writeJSON(w, r, http.StatusNotFound, map[string]string{"error": "not found"})
}

func (api *Api) internalErrorResponse(w http.ResponseWriter, r *http.Request, err ...any) {
	api.Logger.Error("internal server error", err...)
	api.writeJSON(w, r, http.StatusInternalServerError, map[string]string{"error": "something went wrong"})
}
