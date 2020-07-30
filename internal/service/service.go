package service

import (
	"context"
	"encoding/base64"
	"fmt"
	pb "wechat/api"
	"wechat/internal/dao"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)), wire.Bind(new(Server), new(*Service)))

// Service service.
type Service struct {
	ac  *paladin.Map
	dao dao.Dao
}

func (s *Service) GetToken() string {
	return paladin.String(s.ac.Get("token"), "")
}

func (s *Service) GetOriID() string {
	return paladin.String(s.ac.Get("oriID"), "")
}

func (s *Service) GetAppID() string {
	return paladin.String(s.ac.Get("appID"), "")
}

func (s *Service) GetAESKey() []byte {
	paladin.String(s.ac.Get("base64AESKey"), "")

	base64AESKey, err := s.ac.Get("base64AESKey").String()
	if err != nil {
		panic(err)
	}
	if len(base64AESKey) != 43 {
		panic("the length of base64AESKey must equal to 43")
	}
	if aesKey, err := base64.StdEncoding.DecodeString(base64AESKey + "="); err != nil {
		panic(err)
	} else {
		return aesKey
	}
}

func (s *Service) ReplyMessage(ctx context.Context, xmlMsg []byte) (reply []byte, err error) {
	reply = []byte(`<xml><ToUserName><![CDATA[oqeBd0fGbtYTmoVGhHzZ5Nf3-Egc]]></ToUserName>
	<FromUserName><![CDATA[gh_6f7fdb146e4f]]></FromUserName>
	<CreateTime>1596123779</CreateTime>
	<MsgType><![CDATA[text]]></MsgType>
	<Content><![CDATA[a]]></Content>
	<MsgId>22850995649101140</MsgId>
	</xml>`)

	return

}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc demo func.
func (s *Service) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)
	fmt.Printf("hello %s", req.Name)
	return
}

// SayHelloURL bm demo func.
func (s *Service) SayHelloURL(ctx context.Context, req *pb.HelloReq) (reply *pb.HelloResp, err error) {
	reply = &pb.HelloResp{
		Content: "hello " + req.Name,
	}
	fmt.Printf("hello url %s", req.Name)
	return
}

// Ping ping the resource.
func (s *Service) Ping(ctx context.Context, e *empty.Empty) (*empty.Empty, error) {
	return &empty.Empty{}, s.dao.Ping(ctx)
}

// Close close the resource.
func (s *Service) Close() {
}
