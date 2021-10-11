package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetRoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetRoleListLogic {
	return GetRoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRoleListLogic) GetRoleList(req types.SearchRoleRequest) (*types.ApiResponse, error) {
	roles, err := l.svcCtx.RoleModel.List(req.Name)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(roles)), nil
}
