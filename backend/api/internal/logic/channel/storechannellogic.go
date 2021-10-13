package logic

import (
	"context"
	"encoding/json"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

	"go-project/graphics-manage/backend/api/internal/svc"
	types "go-project/graphics-manage/backend/api/internal/types"

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
	audienceConfig, err := json.Marshal(req.Audiences)
	if err != nil {
		return nil, err
	}
	channelInfo := model.GraphicsChannel{
		Name:           req.Name,
		AudienceConfig: string(audienceConfig),
	}
	var channelOwners []model.ChannelOwner
	for _, item := range req.Owners {
		channelOwners = append(channelOwners, model.ChannelOwner{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	channel, err := l.svcCtx.GraphicsChannelModel.UpdateOrCreate(req.Id, channelInfo, channelOwners, l.svcCtx.GraphicsCasbinRuleEnforce)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}

	return (*types.ApiResponse)(helper.ApiSuccess(channel)), nil
}
