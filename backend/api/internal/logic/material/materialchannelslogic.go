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

func (l *MaterialChannelsLogic) MaterialChannels(userId string) (*types.ApiResponse, error) {
	channelIds, err := l.svcCtx.GraphicsChannelModel.GetOwnerChannels(userId, l.svcCtx.GraphicsCasbinRuleEnforce)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	if len(channelIds) == 0 {
		return (*types.ApiResponse)(helper.ApiSuccess(nil)), nil
	}
	channels, err := l.svcCtx.GraphicsChannelModel.ChannelOwner(channelIds)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(channels)), nil
}
