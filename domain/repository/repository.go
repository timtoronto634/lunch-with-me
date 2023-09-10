package repository

import "echo-me/domain/entity/plan"

type PlanRepository interface {
	GetByResearvationID(plan.ReservationID) (plan.Plan, error)
	Store(plan.Plan) error
}
