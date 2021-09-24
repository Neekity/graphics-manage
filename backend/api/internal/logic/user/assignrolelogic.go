package logic

import (
	"context"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AssignRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAssignRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) AssignRoleLogic {
	return AssignRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AssignRoleLogic) AssignRole(req types.AssignRolerRequest) (*types.ApiResponse, error) {
	// todo: add your logic here and delete this line

	return &types.ApiResponse{}, nil
}
