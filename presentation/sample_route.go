package presentation

import (
	"context"
	samplev1 "echo-me/gen/proto/sample/v1"
	"echo-me/gen/proto/sample/v1/samplev1connect"
	"fmt"

	"connectrpc.com/connect"
	"github.com/gorilla/mux"
)

// SampleServiceRPC returns sample service.
type SampleServiceRPC struct{}

func (SampleServiceRPC) Greet(ctx context.Context, req *connect.Request[samplev1.GreetRequest]) (*connect.Response[samplev1.GreetResponse], error) {
	return connect.NewResponse(&samplev1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	}), nil
}

func addSampleServiceRoute(r *mux.Router) {
	path, handler := samplev1connect.NewSampleServiceHandler(SampleServiceRPC{})
	r.PathPrefix(path).Handler(handler)
}
