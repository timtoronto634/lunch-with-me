package presentation

import (
	"context"
	slotv1 "echo-me/presentation/gen/slot/v1"
	"echo-me/presentation/gen/slot/v1/slotv1connect"
	"errors"
	"net/http"

	"connectrpc.com/connect"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

type Server struct {
	e *echo.Echo
}

type slotService struct{}

func (s slotService) SaveSlot(ctx context.Context, req *connect.Request[slotv1.SaveSlotRequest]) (*connect.Response[slotv1.SaveSlotResponse], error) {
	res := connect.NewResponse(&slotv1.SaveSlotResponse{
		Id: 1,
	})
	return res, nil
}

func SetupServer() *Server {
	s := &Server{
		e: echo.New(),
	}
	s.e.Use(sentryecho.New(sentryecho.Options{Repanic: false}))
	slotService := slotService{}
	path, handler := slotv1connect.NewSaveSlotServiceHandler(slotService)

	s.e.POST(path+"SaveSlot", echo.WrapHandler(handler))
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
