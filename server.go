package main

import (
	"database/sql"
	"echo-me/presentation"
	"fmt"
	"log/slog"

	"github.com/getsentry/sentry-go"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		slog.Error(err.Error())
	}
	defer db.Close()

	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://391fa36c79b76bc2db4d04786ff77e3f@o4505783552901120.ingest.sentry.io/4505783556702208",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v", err)
	}

	server := presentation.NewServer()
	server.Start(":1323")
}
