package handler

import (
	"net/http"

	"go-project/graphics-manage/backend/api/internal/logic/role"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetRoleListHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchRoleRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetRoleListLogic(r.Context(), ctx)
		resp, err := l.GetRoleList(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
