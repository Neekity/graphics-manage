package handler

import (
	"net/http"

	"go-project/graphics-manage/backend/api/internal/logic/oauth"
	"go-project/graphics-manage/backend/api/internal/svc"
)

func AttemptLoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewAttemptLoginLogic(r.Context(), ctx)
		url := l.AttemptLogin()
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
}
