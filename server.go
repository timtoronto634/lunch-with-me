package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func saveSlot(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, id+". name:"+name+", email:"+email)
}

func getSlots(c echo.Context) error {
	return c.String(http.StatusOK, "get slots")
}

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

	e := echo.New()
	e.Use(sentryecho.New(sentryecho.Options{Repanic: true}))
	e.PUT("/slots/:id", saveSlot)
	e.GET("/slots", getSlots)
	e.GET("/er", getError)
	e.Logger.Fatal(e.Start(":1323"))
}

func getError(c echo.Context) error {
	return errors.New("mock error")
}
