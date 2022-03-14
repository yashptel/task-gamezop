package controllers

import (
	"net/http"

	"github.com/yashptel/server-b/pkg/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.JSONSuccess(w, http.StatusOK, "OK")
}
