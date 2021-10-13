package logic

import (
	"context"
	"go-project/graphics-manage/backend/common/helper"
	"go-project/graphics-manage/backend/model"

	"go-project/graphics-manage/backend/api/internal/svc"
	"go-project/graphics-manage/backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStoreMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) StoreMenuLogic {
	return StoreMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StoreMenuLogic) StoreMenu(req types.StoreMenuRequest) (*types.ApiResponse, error) {
	menuInfo := model.Menu{
		Name:     req.Name,
		Path:     req.Path,
		Icon:     req.Icon,
		ParentId: req.ParentId,
	}
	menu, err := l.svcCtx.MenuModel.UpdateOrCreate(req.Id, menuInfo)
	if err != nil {
		return (*types.ApiResponse)(helper.ApiError(err.Error(), nil)), nil
	}

	return (*types.ApiResponse)(helper.ApiSuccess(menu)), nil
}
