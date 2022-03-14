package controllers

import (
	"context"
	"net/http"

	"github.com/yashptel/server-b/pkg/config"
	"github.com/yashptel/server-b/pkg/utils"
	"go.uber.org/zap"
)

type RewardController struct {
	ctx  context.Context
	conf config.AppConfig
}

func NewRewardController(ctx context.Context) *RewardController {
	return &RewardController{
		ctx:  ctx,
		conf: config.GetConfig(),
	}
}

func (uc *RewardController) RegisterRewardEvent(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		zap.L().Error("invalid user id")
		utils.JSONError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	zap.L().Info("reward event", zap.String("id", id))

	utils.JSONSuccess(w, http.StatusOK, "reward event registered")
}
