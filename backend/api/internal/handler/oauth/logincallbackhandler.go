package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/api/internal/logic/oauth"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func LoginCallbackHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewLoginCallbackLogic(r.Context(), ctx)
		code := r.URL.Query().Get("code")
		token, err := l.LoginCallback(code)
		if err != nil {
			httpx.Error(w, err)
		} else {
			http.Redirect(w, r, ctx.Config.FrontUrl+token.AccessToken, http.StatusMovedPermanently)
		}
	}
}
