package main

import (
	"github.com/gorilla/mux"
)

func setupApiRoutesV1() (*mux.Router, error) {
	r := mux.NewRouter()
	setupPlantsApiRoutes(r)

	return r.PathPrefix("/api/v1").Subrouter(), nil
}
