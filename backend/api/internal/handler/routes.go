// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	channel "go-project/graphics-manage/backend/api/internal/handler/channel"
	login "go-project/graphics-manage/backend/api/internal/handler/login"
	material "go-project/graphics-manage/backend/api/internal/handler/material"
	menu "go-project/graphics-manage/backend/api/internal/handler/menu"
	message "go-project/graphics-manage/backend/api/internal/handler/message"
	oauth "go-project/graphics-manage/backend/api/internal/handler/oauth"
	permission "go-project/graphics-manage/backend/api/internal/handler/permission"
	role "go-project/graphics-manage/backend/api/internal/handler/role"
	user "go-project/graphics-manage/backend/api/internal/handler/user"
	"go-project/graphics-manage/backend/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/channel/store",
					Handler: channel.StoreChannelHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/detail",
					Handler: channel.ChannelDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/channel/list",
					Handler: channel.ChannelListHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog, serverCtx.CheckDataPermission},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/material/store",
					Handler: material.StoreMaterialHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/delete",
					Handler: material.DeleteMaterialHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/detail",
					Handler: material.MaterialDetailHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/image/upload",
					Handler: material.MaterialImageUploadHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/list",
					Handler: material.MaterialListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/publish",
					Handler: material.MaterialPublishHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/material/channels",
					Handler: material.MaterialChannelsHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/menu/delete",
					Handler: menu.DeleteMenuHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/list",
					Handler: menu.MenuListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/menu/store",
					Handler: menu.StoreMenuHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog, serverCtx.CheckDataPermission},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/Message/change-status",
					Handler: message.MessageChangeStatusHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/Message/list",
					Handler: message.MessageListHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/oauth/login",
				Handler: oauth.AttemptLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/oauth/callback",
				Handler: oauth.LoginCallbackHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/user",
					Handler: user.GetUserListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/role/assign",
					Handler: user.AssignRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/roles",
					Handler: user.UserRolesHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/role",
					Handler: role.GetRoleListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/store",
					Handler: role.StoreRoleHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/role/detail",
					Handler: role.RoleDetailHandler(serverCtx),
				},
			}...,
		),
	)

	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/logout",
				Handler: login.LogoutHandler(serverCtx),
			},
		},
	)

	engine.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth, serverCtx.AccessLog},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/permission",
					Handler: permission.GetPermissionListHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/permission/store",
					Handler: permission.StorePermissionHandler(serverCtx),
				},
				{
					Method:  http.MethodPost,
					Path:    "/permission/detail",
					Handler: permission.PermissionDetailHandler(serverCtx),
				},
			}...,
		),
	)
}
