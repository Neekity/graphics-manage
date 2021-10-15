package handler

import (
	"net/http"

	"go-project/graphics-manage/backend/api/internal/logic/message"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func MessageOwnerListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchMessageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}
		userId := r.Header.Get("UserId")
		l := logic.NewMessageOwnerListLogic(r.Context(), ctx)
		resp, err := l.MessageOwnerList(req, userId)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
