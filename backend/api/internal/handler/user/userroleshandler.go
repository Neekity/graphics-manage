package handler

import (
	"net/http"

	"go-project/graphics-manage/backend/api/internal/logic/user"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func UserRolesHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserRolesRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewUserRolesLogic(r.Context(), ctx)
		resp, err := l.UserRoles(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
