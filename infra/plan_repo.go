package infra

import (
	"context"
	"database/sql"
	"echo-me/domain/entity/plan"
	"echo-me/domain/repository"
)

type PlanRepo struct {
	DB *sql.DB
}

func NewPlanRepo(DB *sql.DB) repository.PlanRepository {
	return &PlanRepo{DB: DB}
}

func (r *PlanRepo) GetByResearvationID(ctx context.Context, resvID plan.ReservationID) (*plan.Plan, error) {
	return &plan.Plan{}, nil
}
func (r *PlanRepo) Store(ctx context.Context, plan *plan.Plan) error {
	return nil
}
