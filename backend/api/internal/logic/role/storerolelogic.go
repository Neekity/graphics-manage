package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreRoleLogic {
	return StoreRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreRoleLogic) StoreRole(req types.StoreAndAssignRolerRequest) (*types.ApiResponse, error) {
	roleInfo := model.Role{CasbinRole: req.CasbinRole, Name: req.Name}
	role, err := l.svcCtx.RoleModel.UpdateOrCreate(req.Id, roleInfo, req.UserIds, l.svcCtx.GraphicsCasbinRuleEnforce)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}

	return (*types.ApiResponse)(helper.ApiSuccess(role)), nil
}
