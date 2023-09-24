package presentation

import (
	"context"
	samplev1 "echo-me/gen/proto/sample/v1"
	"echo-me/gen/proto/sample/v1/samplev1connect"
	"errors"
	"fmt"
	"time"

	"connectrpc.com/connect"
	"github.com/gorilla/mux"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/protobuf/types/known/durationpb"
)

// SampleServiceRPC returns sample service.
type SampleServiceRPC struct{}

func (SampleServiceRPC) Greet(ctx context.Context, req *connect.Request[samplev1.GreetRequest]) (*connect.Response[samplev1.GreetResponse], error) {
	return connect.NewResponse(&samplev1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	}), nil
}

func (SampleServiceRPC) Err(context.Context, *connect.Request[samplev1.ErrRequest]) (*connect.Response[samplev1.ErrResponse], error) {
	return nil, errors.New("sample vanilla error")
}

func (SampleServiceRPC) DetaildErr(context.Context, *connect.Request[samplev1.DetaildErrRequest]) (*connect.Response[samplev1.DetaildErrResponse], error) {
	err := connect.NewError(
		connect.CodeUnavailable,
		errors.New("overloaded: back off and retry"),
	)
	retryInfo := &errdetails.RetryInfo{
		RetryDelay: durationpb.New(10 * time.Second),
	}
	if detail, detailErr := connect.NewErrorDetail(retryInfo); detailErr == nil {
		err.AddDetail(detail)
	}

	return nil, err
}

func addSampleServiceRoute(r *mux.Router) {
	path, handler := samplev1connect.NewSampleServiceHandler(SampleServiceRPC{})
	r.PathPrefix(path).Handler(handler)
}
