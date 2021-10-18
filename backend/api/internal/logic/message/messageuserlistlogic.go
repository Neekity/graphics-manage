package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"

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

func (l *MessageUserListLogic) MessageUserList(req types.SearchMessageRequest, userId string) (*types.ApiResponse, error) {
	messages, err := l.svcCtx.GraphicsMessageModel.List(req.Title, req.ChannelId)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	messages, err = l.svcCtx.GraphicsMessageModel.FilterMessage(l.svcCtx.GraphicsCasbinRuleEnforce, messages, userId, constant.CasbinPermissionRead)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(messages)), nil
}
