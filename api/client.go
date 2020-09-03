package api

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"

	"google.golang.org/grpc"
)

// AppID .
const AppID = "TODO: ADD APP ID"

func NewfileSystemClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (FileSystemClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("direct://default/127.0.0.1:1239"))
	if err != nil {
		return nil, err
	}
	return NewFileSystemClient(cc), nil
}

// direct://default/123.56.29.61:1241
// direct://default/127.0.0.1:1241
func NewtbkClient(cfg *warden.ClientConfig, opts ...grpc.DialOption) (TBKClient, error) {
	client := warden.NewClient(cfg, opts...)
	cc, err := client.Dial(context.Background(), fmt.Sprintf("direct://default/123.56.29.61:1241"))
	if err != nil {
		return nil, err
	}
	return NewTBKClient(cc), nil
}

// 生成 gRPC 代码
//go:generate kratos tool protoc --grpc --bm --swagger api.proto
