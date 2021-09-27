package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DeleteMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) DeleteMaterialLogic {
	return DeleteMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMaterialLogic) DeleteMaterial(req types.DeleteMaterialRequest) (*types.ApiResponse, error) {
	err := l.svcCtx.GraphicsMaterialModel.Delete(req.Id)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(nil)), nil
}
