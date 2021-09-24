package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/material"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func MaterialChannelsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewMaterialChannelsLogic(r.Context(), ctx)
		resp, err := l.MaterialChannels()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
