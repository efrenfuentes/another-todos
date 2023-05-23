package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"

	"github.com/efrenfuentes/todo-backend-go/internal/data"
)

const version = "1.0.0"

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
}

type Application struct {
	Config Config
	Logger *log.Logger
	Models data.Models
}

// The OpenDB() method returns a sql.DB connection pool.
func OpenDB(cfg Config) (*sql.DB, error) {
	// Use sql.Open() to create an empty connection pool, using the DSN from the config
	db, err := sql.Open("postgres", cfg.Db.Dsn)
	if err != nil {
		return nil, err
	}

	// Create the context with a 5-second timeout deadline.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Using PingContext() to establish a new connection to the database, passing in the
	// context we created above as a parameter. If the connection couldn't be established
	// successfully within the 5-second deadline, an error will be returned.
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	// Return the sql.DB connection pool.
	return db, nil
}
