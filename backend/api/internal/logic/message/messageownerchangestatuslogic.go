package logic

import (
	"context"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageOwnerChangeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageOwnerChangeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageOwnerChangeStatusLogic {
	return MessageOwnerChangeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageOwnerChangeStatusLogic) MessageOwnerChangeStatus(req types.MessageOwnerChangeStatusRequest) (*types.ApiResponse, error) {
	// todo: add your logic here and delete this line

	return &types.ApiResponse{}, nil
}
