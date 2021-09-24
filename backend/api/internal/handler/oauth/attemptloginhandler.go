package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/oauth"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func AttemptLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewAttemptLoginLogic(r.Context(), ctx)
		err := l.AttemptLogin()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
