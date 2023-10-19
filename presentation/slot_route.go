package presentation

import (
	"context"
	"echo-me/domain/entity/plan"
	"echo-me/domain/entity/user"
	slotv1 "echo-me/gen/proto/v1"
	"echo-me/gen/proto/v1/slotv1connect"
	"echo-me/infra"
	"fmt"
	"math/rand"
	"time"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type slotService struct{}

func (s slotService) SaveSlot(ctx context.Context, req *connect.Request[slotv1.SaveSlotRequest]) (*connect.Response[slotv1.SaveSlotResponse], error) {
	err := req.Msg.ValidateAll()
	if err != nil {
		return nil, err
	}

	id := int(req.Msg.Id)
	fmt.Println("id", id)
	if id == 0 {
		id = rand.Intn(100)
	}

	res := connect.NewResponse(&slotv1.SaveSlotResponse{
		Id: int64(id),
	})

	repo := infra.NewPlanRepo(infra.DB)
	repo.Save(ctx, plan.NewPlan(user.UserID(uuid.New()), time.Now()))
	return res, nil
}

func addSlotRoute(r *mux.Router) {
	path, handler := slotv1connect.NewSlotServiceHandler(slotService{})
	r.PathPrefix(path).Handler(handler)
}
