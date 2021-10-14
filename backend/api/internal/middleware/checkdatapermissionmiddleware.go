package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/tal-tech/go-zero/rest/httpx"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"io/ioutil"
	"net/http"
)

type CheckDataPermissionMiddleware struct {
	enforcer *casbin.Enforcer
}

type Data struct {
	ID        int `json:"id"`
	ChannelId int `json:"channel_id"`
}

func NewCheckDataPermissionMiddleware(enforcer *casbin.Enforcer) *CheckDataPermissionMiddleware {
	return &CheckDataPermissionMiddleware{enforcer: enforcer}
}

func (m *CheckDataPermissionMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Header.Get("UserId")
		if r.URL.Path == "/material/channels" {
			next(w, r)
			return
		}
		if err := m.enforcer.LoadPolicy(); err != nil {
			panic(err)
		}
		bodyInfo, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logx.Error(err.Error())
		}
		err = r.Body.Close()
		if err != nil {
			httpx.OkJson(w, helper.ApiError(err.Error(), nil))
			return
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyInfo))

		var data Data
		if err := json.Unmarshal(bodyInfo, &data); err != nil {
			httpx.OkJson(w, helper.ApiError(err.Error(), nil))
		}

		if flag, _ := m.enforcer.Enforce(userId, fmt.Sprintf(constant.CasbinMaterialPolicy, data.ID), constant.CasbinPermissionWrite); flag {
			next(w, r)
			return
		}
		if data.ChannelId != 0 {
			if flag := m.enforcer.HasGroupingPolicy(userId, fmt.Sprintf(constant.CasbinChannelRole, data.ChannelId)); flag {
				next(w, r)
				return
			}
		}
		httpx.OkJson(w, helper.ApiError("您没有该资源的权限！", nil))
	}
}
