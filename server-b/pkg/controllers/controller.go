package controllers

import (
	"context"

	"github.com/gorilla/mux"
)

const prefix = "/api"

func NewRouter(ctx context.Context) *mux.Router {
	router := mux.NewRouter()
	router = router.PathPrefix(prefix).Subrouter()

	SetupHealthCheck(router)
	SetupReward(ctx, router)
	return router
}

func SetupHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods("GET")
}

func SetupReward(ctx context.Context, router *mux.Router) {
	uc := NewRewardController(ctx)
	router.HandleFunc("/reward", uc.RegisterRewardEvent).Methods("GET")
}
