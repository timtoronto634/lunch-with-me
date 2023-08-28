package main

import (
	"database/sql"
	"net/http"

	"slog"

	_ "github.com/mattn/go-sqlite3"

	"github.com/labstack/echo/v4"
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
		slog.Error(err)
	}
	defer db.Close()

	e := echo.New()
	e.PUT("/slots/:id", saveSlot)
	e.GET("/slots", getSlots)
	e.Logger.Fatal(e.Start(":1323"))
}
