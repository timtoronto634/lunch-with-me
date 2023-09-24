package presentation

import (
	"context"
	slotv1 "echo-me/gen/proto/v1"
	"echo-me/gen/proto/v1/slotv1connect"

	"connectrpc.com/connect"
	"github.com/gorilla/mux"
)

type slotService struct{}

func (s slotService) SaveSlot(ctx context.Context, req *connect.Request[slotv1.SaveSlotRequest]) (*connect.Response[slotv1.SaveSlotResponse], error) {
	err := req.Msg.ValidateAll()
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&slotv1.SaveSlotResponse{
		Id: 1,
	})
	return res, nil
}

func addSlotRoute(r *mux.Router) {
	path, handler := slotv1connect.NewSlotServiceHandler(slotService{})
	r.PathPrefix(path).Handler(handler)
}
