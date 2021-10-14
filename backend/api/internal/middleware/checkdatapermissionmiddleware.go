package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"net/http"
)

type CheckDataPermissionMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewCheckDataPermissionMiddleware(enforcer *casbin.Enforcer) *CheckDataPermissionMiddleware {
	return &CheckDataPermissionMiddleware{enforcer: enforcer}
}

func (m *CheckDataPermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("UserId")
		if err := m.enforcer.LoadPolicy(); err != nil {
			panic(err)
		}
		flag, err := m.enforcer.Enforce(userId, r.URL.Path, constant.CasbinPermissionWrite)
		if err != nil {
			panic(err)
		}
		if !flag {
			httpx.OkJson(w, helper.ApiError("未拥有该路由的权限", nil))
			return
		}
		next(w, r)
	}
}
