package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageUserChannelsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageUserChannelsLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageUserChannelsLogic {
	return MessageUserChannelsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageUserChannelsLogic) MessageUserChannels(userId string) (*types.ApiResponse, error) {
	channelIds, err := l.svcCtx.GraphicsChannelModel.GetUserChannels(userId)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	if len(channelIds) == 0 {
		return (*types.ApiResponse)(helper.ApiSuccess(nil)), nil
	}
	channels, err := l.svcCtx.GraphicsChannelModel.GetChannels(channelIds)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(channels)), nil
}
