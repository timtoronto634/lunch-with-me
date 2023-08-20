package main

import (
	"net/http"

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
	e := echo.New()
	e.PUT("/slots/:id", saveSlot)
	e.GET("/slots", getSlots)
	e.Logger.Fatal(e.Start(":1323"))
}
