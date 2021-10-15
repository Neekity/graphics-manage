package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageOwnerDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageOwnerDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageOwnerDetailLogic {
	return MessageOwnerDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageOwnerDetailLogic) MessageOwnerDetail(req types.MessageDetailRequest) (*types.ApiResponse, error) {
	message, err := l.svcCtx.GraphicsMessageModel.OwnerDetail(req.Id)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(message)), nil
}
