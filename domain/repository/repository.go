package repository

func UserRepository() error {
	return nil
}

type PlanRepositoryInterface interface {
	Save()
}
