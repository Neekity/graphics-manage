package logic

import (
	"context"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageUserListLogic {
	return MessageUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageUserListLogic) MessageUserList(req types.SearchMessageRequest) (*types.ApiResponse, error) {
	// todo: add your logic here and delete this line

	return &types.ApiResponse{}, nil
}
