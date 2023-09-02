package infra

import (
	"echo-me/domain/entity/plan"

	"github.com/jackc/pgx/v5"
)

type PlanRepo struct {
	conn pgx.QueryExecMode
}

func (r *PlanRepo) Save(plan plan.Plan) error {
	return nil
}
func (r *PlanRepo) GetByResearvationID(reservationID plan.ReservationID) plan.Plan {
	return plan.Plan{}
}
