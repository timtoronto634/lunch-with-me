package infra

import (
	"context"
	"database/sql"
	"echo-me/domain/entity/plan"
)

type PlanRepo struct {
	DB *sql.DB
}

func NewPlanRepo(DB *sql.DB) plan.Repository {
	return &PlanRepo{DB: DB}
}

func (r *PlanRepo) GetByResearvationID(ctx context.Context, resvID plan.ReservationID) (*plan.Plan, error) {
	return &plan.Plan{}, nil
}
func (r *PlanRepo) Save(ctx context.Context, plan *plan.Plan) error {
	return nil
}
