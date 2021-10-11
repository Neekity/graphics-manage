package handler

import (
	"net/http"
	"strconv"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/menu"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func MenuListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewMenuListLogic(r.Context(), ctx)
		userId, err := strconv.Atoi(r.Header.Get("UserId"))
		resp, err := l.MenuList(uint(userId))
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
