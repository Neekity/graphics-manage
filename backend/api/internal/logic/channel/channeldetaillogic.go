package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChannelDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChannelDetailLogic {
	return ChannelDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelDetailLogic) ChannelDetail(req types.ChanenlDetailRequest) (*types.ApiResponse, error) {
	data, err := l.svcCtx.GraphicsChannelModel.FindOne(req.Id)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(data)), nil
}
