package service

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"time"
	pb "wechat/api"
	"wechat/internal/dao"
	"wechat/model"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/pkg/log"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)), wire.Bind(new(Server), new(*Service)))
var aesKey string

// Service service.
type Service struct {
	ac               *paladin.Map
	dao              dao.Dao
	fileSystemClient pb.FileSystemClient
}

func (s *Service) MediaIDGet(ctx context.Context, req model.MediaIDReq, args ...interface{}) (resp model.MediaIDResp, err error) {
	option := make([]grpc.CallOption, len(args))
	for i := range args {
		if o, ok := args[i].(grpc.CallOption); ok {
			option[i] = o
		}
	}
	get := new(pb.MediaIDResp)
	get, err = s.fileSystemClient.MediaIDGet(ctx, &pb.MediaIDReq{
		FakeID:    req.FakeID,
		Timestamp: req.Timestamp,
	}, option...)
	if err != nil {
		err = fmt.Errorf("MediaIDGet: (%w)", err)
		return
	}
	if get == nil {
		err = errors.New("MediaIDGet: rpc get is nil")
		return
	}
	resp.MediaID = get.MediaID
	return
}

func (s *Service) GetToken() string {
	return paladin.String(s.ac.Get("token"), "")
}

func (s *Service) GetSecret() string {
	return paladin.String(s.ac.Get("secret"), "")
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
	rx, err := model.FromEnvelope(xmlMsg)
	if err != nil {
		err = fmt.Errorf("ReplyMessage: (%w)", err)
		return
	}
	log.Info("receive message: %s", rx)
	var txMessage model.TxMessageKind

	switch rx.MsgType {
	case model.RxMessageTypeText:
		// 不用ok是因为var _ RxTextMessageExtra = (*rxTextMessageSpecifics)(nil)检查过了
		text, _ := rx.Text()
		// 尴尬了.
		if text.GetContent() == "【收到不支持的消息类型，暂无法显示】" {
			mediaIDResp, err := s.MediaIDGet(context.Background(), model.MediaIDReq{
				FakeID:    rx.FromUserName,
				Timestamp: rx.CreateTime,
			})
			if err != nil {
				err = fmt.Errorf("ReplyMessage: (%w)", err)
			}
			txMessage = model.NewTxImageMessageSpecifics(mediaIDResp.MediaID)
			break
		}
		txMessage = s.handleText(text)
	case model.RxMessageTypeImage:
		image, _ := rx.Image()
		txMessage = s.handleImage(image)
	case model.RxMessageTypeVoice:
		voice, _ := rx.Voice()
		txMessage = s.handleVoice(voice)
	case model.RxMessageTypeVideo:
		video, _ := rx.Video()
		txMessage = s.handleVideo(video)
	case model.RxMessageTypeLocation:
		location, _ := rx.Location()
		txMessage = s.handleLocation(location)
	case model.RxMessageTypeLink:
		link, _ := rx.Link()
		txMessage = s.handleLink(link)
	case model.RxMessageTypeShortVideo:
		video, _ := rx.ShortVideo()
		txMessage = s.handleShortVideo(video)
	default:
		err = fmt.Errorf("ReplyMessage: can not handle message '%#v'", rx)
		return
	}

	if txMessage == nil {
		err = fmt.Errorf("ReplyMessage: Unimplemented interface")
		specifics := model.NewTxTextMessageSpecifics("不知道你在说啥")
		return xml.Marshal(model.NewTxMessage(rx.FromUserName, rx.ToUserName, specifics.TxMessageType(), time.Now().Unix(), rx.MsgID, specifics))
	}

	tx := model.NewTxMessage(rx.FromUserName, rx.ToUserName, txMessage.TxMessageType(), time.Now().Unix(), rx.MsgID, txMessage)
	log.Info("reply message: %s", tx)
	return xml.Marshal(tx)

}
func (s *Service) handleText(text model.RxTextMessageExtra) model.TxMessageKind {
	switch content := text.GetContent(); content {
	case "我要表情包":
		return model.NewTxTextMessageSpecifics("不给")
	}
	return nil
}
func (s *Service) handleImage(image model.RxImageMessageExtra) model.TxMessageKind {
	return nil
}
func (s *Service) handleVoice(voice model.RxVoiceMessageExtra) model.TxMessageKind {
	return nil

}
func (s *Service) handleVideo(video model.RxVideoMessageExtra) model.TxMessageKind {
	return nil

}
func (s *Service) handleLocation(location model.RxLocationMessageExtra) model.TxMessageKind {
	return nil

}
func (s *Service) handleLink(link model.RxLinkMessageExtra) model.TxMessageKind {
	return nil

}
func (s *Service) handleShortVideo(image model.RxShortVideoMessageExtra) model.TxMessageKind {
	return nil

}

// New new a service and return.
func New(d dao.Dao) (s *Service, cf func(), err error) {
	s = &Service{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	if s.fileSystemClient, err = pb.NewClient(&warden.ClientConfig{
		Timeout: xtime.Duration(time.Second * 5),
	}); err != nil {
		panic(err)
	}
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
