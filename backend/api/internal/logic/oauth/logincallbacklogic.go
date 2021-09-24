package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go-project/graphics-manage/backend/api/internal/svc"
)

type LoginCallbackLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginCallbackLogic {
	return LoginCallbackLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginCallbackLogic) LoginCallback() error {
	// todo: add your logic here and delete this line

	return nil
}
