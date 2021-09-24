package logic

import (
	"context"

	"github.com/tal-tech/go-zero/core/logx"
	"go-project/graphics-manage/backend/api/internal/svc"
)

type AttemptLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAttemptLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) AttemptLoginLogic {
	return AttemptLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AttemptLoginLogic) AttemptLogin() error {
	// todo: add your logic here and delete this line

	return nil
}
