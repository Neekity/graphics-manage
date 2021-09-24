package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/constant"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MaterialImageUploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialImageUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) MaterialImageUploadLogic {
	return MaterialImageUploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialImageUploadLogic) MaterialImageUpload(channelId int, name string, url string) (*types.ApiResponse, error) {
	materialImage := model.GraphicsMaterial{
		Name:      name,
		ChannelID: channelId,
		URL:       url,
		Type:      constant.MaterialImage,
	}
	material, err := l.svcCtx.GraphicsMaterialModel.UpdateOrCreate(0, materialImage)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}
	return (*types.ApiResponse)(helper.ApiSuccess(material)), nil
}
