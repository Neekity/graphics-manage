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
	"strings"
)

type CheckDataPermissionMiddleware struct {
	enforcer *casbin.Enforcer
}

type Data struct {
	ID        int    `json:"id"`
	ChannelId int    `json:"channel_id"`
	DataType  string `json:"data_type"`
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

		dataType := constant.CasbinPermissionWrite
		if len(data.DataType) > 0 && data.DataType == constant.CasbinPermissionRead {
			dataType = constant.CasbinPermissionRead
		}

		if data.ID != 0 {
			dataPermission := fmt.Sprintf(constant.CasbinMaterialPolicy, data.ID)
			if strings.Contains(r.URL.Path, "message") {
				dataPermission = fmt.Sprintf(constant.CasbinMessagePolicy, data.ID)
			}

			if flag, _ := m.enforcer.Enforce(userId, dataPermission, dataType); flag {
				next(w, r)
				return
			}
		} else {
			if dataType == constant.CasbinPermissionRead {
				next(w, r)
				return
			}
			if data.ChannelId != 0 {
				if flag := m.enforcer.HasGroupingPolicy(userId, fmt.Sprintf(constant.CasbinChannelRole, data.ChannelId)); flag {
					next(w, r)
					return
				}
			} else {
				policy := m.enforcer.GetFilteredNamedGroupingPolicy(constant.CasbinDefaultRole,
					0, userId)
				for _, item := range policy {
					role := strings.Split(item[1], constant.CasbinPermissionSymbol)
					if len(role) != 3 || !strings.Contains(item[1], constant.CasbinChannelRole[:len(constant.CasbinChannelRole)-2]) {
						continue
					}
					next(w, r)
					return
				}
			}
		}

		httpx.OkJson(w, helper.ApiError("您没有该资源的权限！", nil))
	}
}
