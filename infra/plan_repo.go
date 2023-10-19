package infra

import (
	"context"
	"echo-me/domain/entity/plan"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PlanRepo struct {
	DB *pgxpool.Pool
}

func NewPlanRepo(DB *pgxpool.Pool) plan.Repository {
	return &PlanRepo{DB: DB}
}

func (r *PlanRepo) GetByResearvationID(ctx context.Context, resvID plan.ReservationID) (*plan.Plan, error) {
	return &plan.Plan{}, nil
}
func (r *PlanRepo) Save(ctx context.Context, plan *plan.Plan) error {
	return nil
}
