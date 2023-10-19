package main

import (
	"context"
	"echo-me/infra"
	"echo-me/presentation"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	connStr := "postgresql://postgres:postgres@localhost:5432/psql?sslmode=disable"
	db, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}
	infra.DB = db

	router := presentation.SetupServer(db)
	loggedRouter := handlers.LoggingHandler(os.Stdout, router)
	port := 8088
	slog.Info(fmt.Sprintf("running server at %v", port))
	slog.Error("error: %v", http.ListenAndServe(
		fmt.Sprintf("localhost:%d", port),
		h2c.NewHandler(loggedRouter, &http2.Server{}),
	))
}
