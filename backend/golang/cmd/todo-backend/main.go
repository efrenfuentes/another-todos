package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/efrenfuentes/todo-backend-go/internal/app"
	"github.com/efrenfuentes/todo-backend-go/internal/data"
)

func main() {
	var cfg app.Config

	// Read the value of the port and env command-line flags into the config struct. We
	// default to using the port number 4000 and the environment "development" if no
	// corresponding flags are provided
	flag.IntVar(&cfg.Port, "port", 4000, "API server port")
	flag.StringVar(&cfg.Env, "env", "development", "Environment (development|staging|production)")

	// Read the DSN value from db-dsn command-line flag into the config struct. We default
	// to using the develoipment DSN "postgres://postgres:postgres@localhost/todos_dev?sslmode=disable" if no corresponding flag is provided.
	flag.StringVar(&cfg.Db.Dsn, "db-dsn", "postgres://postgres:postgres@localhost/todos_dev?sslmode=disable", "Postgres DSN")
	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream,
	// prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := app.OpenDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	logger.Printf("database connection pool established")

	app := &app.Application{
		Config: cfg,
		Logger: logger,
		Models: data.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		Handler:      app.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.Env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)
}
