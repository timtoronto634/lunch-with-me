package presentation

import (
	"context"
	"echo-me/gen/proto/sample/v1/samplev1connect"
	v1 "echo-me/gen/proto/v1"
	"echo-me/gen/proto/v1/slotv1connect"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

type slotService struct{}

func (s slotService) SaveSlot(ctx context.Context, req *connect.Request[v1.SaveSlotRequest]) (*connect.Response[v1.SaveSlotResponse], error) {
	err := req.Msg.ValidateAll()
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&v1.SaveSlotResponse{
		Id: 1,
	})
	return res, nil
}

func SetupServer(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()
	// r.Handle("/", LoggingMiddleware)
	r.Handle(slotv1connect.NewSlotServiceHandler(slotService{}))
	r.Handle(samplev1connect.NewSampleServiceHandler(UnimplementedSampleServiceHandler{}))
	return r
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Start time of the request
		start := time.Now()

		// Serve the request to the next middleware or handler
		next.ServeHTTP(w, r)

		// Calculate the request duration
		duration := time.Since(start)

		// Log the request details
		log.Printf("%s %s %v", r.Method, r.URL.Path, duration)
	})
}
