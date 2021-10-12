package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/login"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func LogoutHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLogoutLogic(r.Context(), ctx)
		resp, err := l.Logout()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
