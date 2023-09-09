package main

import (
	"database/sql"
	"echo-me/presentation"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		slog.Error(err.Error())
	}
	defer db.Close()

	server := presentation.SetupServer()
	server.Start(":1323")
}
