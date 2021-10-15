package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageUserDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageUserDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageUserDetailLogic {
	return MessageUserDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageUserDetailLogic) MessageUserDetail(req types.MessageDetailRequest) (*types.ApiResponse, error) {
	message, err := l.svcCtx.GraphicsMessageModel.UserDetail(req.Id)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(message)), nil
}
