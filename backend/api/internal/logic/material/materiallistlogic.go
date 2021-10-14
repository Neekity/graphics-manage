package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MaterialListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MaterialListLogic {
	return MaterialListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialListLogic) MaterialList(req types.SearchMaterialRequest, userId string) (*types.ApiResponse, error) {
	materials, err := l.svcCtx.GraphicsMaterialModel.List(req.Name, req.Type, req.ChannelId)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	materials, err = l.svcCtx.GraphicsMaterialModel.FilterMaterial(l.svcCtx.GraphicsCasbinRuleEnforce, materials, userId)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(materials)), nil
}
