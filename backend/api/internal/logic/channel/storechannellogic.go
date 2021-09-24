package logic

import (
	"context"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreChannelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreChannelLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreChannelLogic {
	return StoreChannelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreChannelLogic) StoreChannel(req types.StoreChannelRequest) (*types.ApiResponse, error) {
	// todo: add your logic here and delete this line

	return &types.ApiResponse{}, nil
}
