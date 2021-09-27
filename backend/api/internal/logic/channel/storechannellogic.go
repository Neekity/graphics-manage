package logic

import (
	"context"
	"encoding/json"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

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
	audienceConfig, _ := json.Marshal(req.Audiences)

	channel := model.GraphicsChannel{
		Name:           req.Name,
		AudienceConfig: string(audienceConfig),
	}
	material, err := l.svcCtx.GraphicsChannelModel.UpdateOrCreate(req.Id, channel)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(material)), nil
}
