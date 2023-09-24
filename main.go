package main

import (
	"context"
	"echo-me/presentation"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	connStr := "postgresql://postgres:devsample@localhost:5432/psql?sslmode=disable"
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	router := presentation.SetupServer(pool)
	http.ListenAndServe(
		"localhost:8080",
		h2c.NewHandler(router, &http2.Server{}),
	)
}
