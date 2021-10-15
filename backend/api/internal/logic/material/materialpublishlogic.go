package logic

import (
	"context"
	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type MaterialPublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) MaterialPublishLogic {
	return MaterialPublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialPublishLogic) MaterialPublish(req types.PublishMaterialRequest, userId int) (*types.ApiResponse, error) {
	material, err := l.svcCtx.GraphicsMaterialModel.FindOne(req.Id)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	sendTime, err := time.Parse("2006-01-02 15:04:05", req.SendTime)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	messageInfo := model.GraphicsMessage{
		SenderUserID:    userId,
		Type:            material.Type,
		Title:           material.Name,
		Abstract:        material.Abstract,
		CoverPictureURL: material.URL,
		ChannelID:       material.ChannelID,
		Receivers:       req.Receivers,
		Content:         material.Content,
		Status:          0,
		SendTime:        helper.MyTime{Time: sendTime},
		Author:          req.Author,
	}
	message, err := l.svcCtx.GraphicsMessageModel.Create(messageInfo)
	//TODO 推送消息至用户
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(message)), nil
}
