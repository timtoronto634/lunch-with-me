package presentation

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupServer(db *pgxpool.Pool) *mux.Router {
	r := mux.NewRouter()
	addSlotRoute(r)
	addSampleServiceRoute(r)
	return r
}
