package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MaterialChannelsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialChannelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) MaterialChannelsLogic {
	return MaterialChannelsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialChannelsLogic) MaterialChannels() (*types.ApiResponse, error) {
	//channelIds, err := l.svcCtx.GraphicsCasbinRuleModel.GetOwnerChannels(1)
	//if err != nil {
	//	return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	//}
	channels, err := l.svcCtx.GraphicsChannelModel.ChannelOwner([]int{1, 2, 3, 4, 5, 6})
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(channels)), nil
}
