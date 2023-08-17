package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/mrtj458/bolttodo/pkg/api"
	"github.com/mrtj458/bolttodo/pkg/db"
)

func main() {
	portFlag := flag.Int("port", 3000, "port to be used by the API")
	dbNameFlag := flag.String("db", "todos.db", "filename for the bolt database")
	debugFlag := flag.Bool("debug", false, "show debug level logs")
	flag.Parse()

	logOpts := &slog.HandlerOptions{}
	if *debugFlag {
		logOpts.Level = slog.LevelDebug
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, logOpts))

	bdb, err := db.Open(*dbNameFlag)
	if err != nil {
		panic(err)
	}
	defer bdb.Close()

	api := api.Api{
		Debug:       *debugFlag,
		Port:        *portFlag,
		DB:          bdb,
		Logger:      logger,
		TodoService: &db.TodoService{DB: bdb},
	}

	err = api.Start()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
