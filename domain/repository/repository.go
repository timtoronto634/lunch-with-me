package repository

import "echo-me/domain/entity/plan"

func UserRepository() error {
	return nil
}

type PlanRepositoryInterface interface {
	GetByResearvationID(plan.ReservationID) plan.Plan
	Save(plan.Plan) error
}
