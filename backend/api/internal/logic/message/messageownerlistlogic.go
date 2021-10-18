package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MessageOwnerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageOwnerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MessageOwnerListLogic {
	return MessageOwnerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageOwnerListLogic) MessageOwnerList(req types.SearchMessageRequest, userId string) (*types.ApiResponse, error) {
	messages, err := l.svcCtx.GraphicsMessageModel.List(req.Title, req.ChannelId)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	messages, err = l.svcCtx.GraphicsMessageModel.FilterMessage(l.svcCtx.GraphicsCasbinRuleEnforce, messages, userId, constant.CasbinPermissionWrite)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(messages)), nil
}
