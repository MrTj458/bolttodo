package api

import (
	"fmt"
	"net/http"

	"log/slog"

	"github.com/mrtj458/bolttodo/pkg/db"
	"go.etcd.io/bbolt"
)

type Api struct {
	Port   int
	DB     *bbolt.DB
	Logger *slog.Logger

	TodoService *db.TodoService
}

func (api *Api) Start() error {
	router := api.routes()

	api.Logger.Info("Starting server", "port", api.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", api.Port), router)
}
