package service

import (
	"context"
	pb "wechat/api"
)

type Server interface {
	pb.DemoServer
	WxService
}

// 微信公众号服务
type WxService interface {
	GetToken() string
	GetOriID() string
	GetAppID() string
	GetAESKey() []byte
	ReplyMessage(ctx context.Context, xmlMsg []byte) (reply []byte, err error)
}
