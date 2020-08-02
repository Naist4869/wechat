package service

import (
	"context"
	pb "wechat/api"
	"wechat/model"
)

type Server interface {
	pb.DemoServer
	WxService
}

// 微信公众号服务
type WxService interface {
	GetToken() string
	GetSecret() string
	GetOriID() string
	GetAppID() string
	GetAESKey() []byte
	ReplyMessage(ctx context.Context, xmlMsg []byte) (reply []byte, err error)
}

// 文件系统服务
type FileSystemService interface {
	MediaIDGet(ctx context.Context, req model.MediaIDReq, args ...interface{}) (resp model.MediaIDResp, err error)
}
