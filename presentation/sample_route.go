package presentation

import (
	"context"
	samplev1 "echo-me/gen/proto/sample/v1"
	"errors"

	"connectrpc.com/connect"
)

// UnimplementedSampleServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSampleServiceHandler struct{}

func (UnimplementedSampleServiceHandler) Greet(context.Context, *connect.Request[samplev1.GreetRequest]) (*connect.Response[samplev1.GreetResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("proto.sample.v1.SampleService.Greet is not implemented"))
}
