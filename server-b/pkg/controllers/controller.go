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
	SetupUser(ctx, router)
	return router
}

func SetupHealthCheck(router *mux.Router) {
	router.HandleFunc("/health", HealthCheck).Methods("GET")
}

func SetupUser(ctx context.Context, router *mux.Router) {
	uc := NewUserController(ctx)
	router.HandleFunc("/user", uc.RegisterUserEvent).Methods("GET")
}
