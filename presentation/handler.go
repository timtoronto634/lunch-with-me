package presentation

import (
	"context"
	slotv1 "echo-me/gen/proto/v1"
	"echo-me/gen/proto/v1/slotv1connect"
	"fmt"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
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
	s.e.Use(LoggingMiddleware)
	path, handler := slotv1connect.NewSlotServiceHandler(slotService{})
	fmt.Println(path, handler)
	s.e.POST(path, echo.WrapHandler(handler))
	s.e.GET("/slots", getSlots)
	return s
}

func (s *Server) Start(addr string) {
	s.e.Start(addr)
}

func getSlots(c echo.Context) error {
	return c.String(http.StatusOK, "get slots")
}

// func LoggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Start time of the request
// 		start := time.Now()

// 		// Serve the request to the next middleware or handler
// 		next.ServeHTTP(w, r)

// 		// Calculate the request duration
// 		duration := time.Since(start)

// 		// Log the request details
// 		log.Printf("%s %s %v", r.Method, r.URL.Path, duration)
// 	})
// }

// LoggingMiddleware logs request details such as method, path, and duration
func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Start time of the request
		start := time.Now()

		// Serve the request to the next middleware or handler
		err := next(c)

		// Calculate the request duration
		duration := time.Since(start)

		// Log the request details
		log.Printf("%s %s %v", c.Request().Method, c.Request().URL.Path, duration)

		return err
	}
}
