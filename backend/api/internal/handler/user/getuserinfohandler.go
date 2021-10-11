package handler

import (
	"net/http"
	"strconv"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/user"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func GetUserInfoHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewGetUserInfoLogic(r.Context(), ctx)
		userId, err := strconv.Atoi(r.Header.Get("UserId"))
		resp, err := l.GetUserInfo(uint(userId))

		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
