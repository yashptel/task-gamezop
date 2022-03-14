package controllers

import (
	"context"
	"net/http"

	"github.com/yashptel/go-api-template/pkg/config"
	"github.com/yashptel/go-api-template/pkg/debounce"
	"github.com/yashptel/go-api-template/pkg/utils"
	"go.uber.org/zap"
)

type UserController struct {
	ctx            context.Context
	debounceClient debounce.DebounceClient
	conf           config.AppConfig
}

func NewUserController(ctx context.Context) *UserController {
	return &UserController{
		ctx:            ctx,
		debounceClient: debounce.NewDebounce(ctx),
		conf:           config.GetConfig(),
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

	err := uc.debounceClient.Set(uc.ctx, id, true, uc.conf.DebounceTime)
	if err != nil {
		zap.L().Error("redis set error", zap.Error(err))
		utils.JSONError(w, http.StatusInternalServerError, "please try again later")
		return
	}
	utils.JSONSuccess(w, http.StatusOK, "user event registered")
}
