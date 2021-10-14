package handler

import (
	"net/http"

	"go-project/graphics-manage/backend/api/internal/logic/material"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MaterialListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchMaterialRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewMaterialListLogic(r.Context(), ctx)
		userId := r.Header.Get("UserId")
		resp, err := l.MaterialList(req, userId)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
