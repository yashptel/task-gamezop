package controllers

import (
	"context"
	"net/http"

	"github.com/yashptel/server-b/pkg/config"
	"github.com/yashptel/server-b/pkg/utils"
	"go.uber.org/zap"
)

type UserController struct {
	ctx  context.Context
	conf config.AppConfig
}

func NewUserController(ctx context.Context) *UserController {
	return &UserController{
		ctx:  ctx,
		conf: config.GetConfig(),
	}
}

func (uc *UserController) RegisterUserEvent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		zap.L().Error("invalid user id")
		utils.JSONError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	zap.L().Info("user event", zap.String("id", id))

	utils.JSONSuccess(w, http.StatusOK, "user event registered")
}
