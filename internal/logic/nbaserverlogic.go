package logic

import (
	"context"

	"nbaserver/internal/svc"
	"nbaserver/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type NbaserverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewNbaserverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NbaserverLogic {
	return &NbaserverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *NbaserverLogic) Nbaserver(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
