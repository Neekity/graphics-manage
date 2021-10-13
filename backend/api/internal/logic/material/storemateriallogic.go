package logic

import (
	"context"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreMaterialLogic {
	return StoreMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreMaterialLogic) StoreMaterial(req types.StoreMaterialRequest) (*types.ApiResponse, error) {
	material := model.GraphicsMaterial{
		Type:      req.Type,
		Name:      req.Name,
		Content:   req.Content,
		ChannelID: req.ChannelId,
		URL:       req.Url,
		Abstract:  req.Abstract,
	}

	menu, err := l.svcCtx.GraphicsMaterialModel.UpdateOrCreate(req.Id, material)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}

	return (*types.ApiResponse)(helper.ApiSuccess(menu)), nil
}
