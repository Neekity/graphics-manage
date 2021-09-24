package logic

import (
	"context"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ChannelListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChannelListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChannelListLogic {
	return ChannelListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChannelListLogic) ChannelList(req types.SearchChannelRequest) (*types.ApiResponse, error) {
	// todo: add your logic here and delete this line

	return &types.ApiResponse{}, nil
}
