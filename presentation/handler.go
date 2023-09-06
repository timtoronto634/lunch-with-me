package presentation

import (
	"errors"
	"net/http"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

func SetupServer() *Server {
	s := &Server{
		e: echo.New(),
	}
	s.e.Use(sentryecho.New(sentryecho.Options{Repanic: false}))
	s.e.PUT("/slots/:id", saveSlot)
	s.e.GET("/slots", getSlots)
	s.e.GET("/er", getError)
	return s
}

func (s *Server) Start(addr string) {
	s.e.Start(addr)
}

func saveSlot(c echo.Context) error {
	id := c.Param("id")
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, id+". name:"+name+", email:"+email)
}

func getSlots(c echo.Context) error {
	return c.String(http.StatusOK, "get slots")
}

func getError(c echo.Context) error {
	return errors.New("mock error")
}
