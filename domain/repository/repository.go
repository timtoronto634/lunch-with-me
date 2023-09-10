package repository

import (
	"context"
	"echo-me/domain/entity/plan"
)

type PlanRepository interface {
	GetByResearvationID(context.Context, plan.ReservationID) (*plan.Plan, error)
	Store(context.Context, *plan.Plan) error
}
