package service

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync"
	"time"
	pb "wechat/api"
	"wechat/internal/dao"
	"wechat/model"

	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"

	"github.com/go-kratos/kratos/pkg/net/rpc/warden"
	xtime "github.com/go-kratos/kratos/pkg/time"

	"google.golang.org/grpc"

	"github.com/go-kratos/kratos/pkg/log"

	"github.com/go-kratos/kratos/pkg/conf/paladin"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.DemoServer), new(*Service)), wire.Bind(new(pb.WechatServer), new(*Service)), wire.Bind(new(Server), new(*Service)))
var tklRe = regexp.MustCompile(`([￥$€₤₳¢¤฿฿₵₡₫ƒ₲₭£₥₦₱〒₮₩₴₪៛﷼₢M₰₯₠₣₧ƒ][a-zA-Z0-9]{9,11}[￥$€₤₳¢¤฿฿₵₡₫ƒ₲₭£₥₦₱〒₮₩₴₪៛﷼₢M₰₯₠₣₧ƒ])`)

// Service service.
type Service struct {
	ac               *paladin.Map
	dao              dao.Dao
	client           *bm.Client
	once             sync.Once
	fileSystemClient pb.FileSystemClient
	tbkClient        pb.TBKClient
	accessToken      string
}

func (s *Service) BalanceTemplateMsgSend(ctx context.Context, req *pb.BalanceTemplateMsgSendReq) (resp *empty.Empty, err error) {
	query := url.Values{}
	query.Set("access_token", s.AccessToken())
	msg := model.TemplateMsg{
		Touser:     req.UserID,
		TemplateID: "64khik56ODNd-4mfTFSt8D-CTF963ObT_0CtzETGY9I",
		URL:        "",
		Data: model.BalanceTemplateMsgData{
			Title: model.TemplateMsgFirst{
				Value: fmt.Sprintf("尊敬的用户: \n您的订单: %s\n%s\n已经结算完成, 发送提现即可0元起提哦", req.OrderID, req.Title),
				Color: "#4682b4",
			},
			EarningTime: model.TemplateMsgKeyword1{
				Value: req.EarningTime,
				Color: "#4682b4",
			},
			Salary: model.TemplateMsgKeyword2{
				Value: req.Salary,
				Color: "#4682b4",
			},
			Balance: model.TemplateMsgKeyword3{
				Value: req.Balance,
				Color: "#4682b4",
			},
			Remark: model.TemplateMsgRemark{
				Value: "12期全免特权卡+善诊双人体验, 99元立抢>>",
				Color: "#FE4365",
			},
		},
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return
	}
	v := new(struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Msgid   int    `json:"msgid"`
	})
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", s.AccessToken()), bytes.NewReader(marshal))
	if err != nil {
		return
	}
	err = s.client.Do(ctx, request, v)
	if err != nil {
		return
	}
	if v.Errcode != 0 {
		err = fmt.Errorf("错误代码: %d,错误信息: %s", v.Errcode, v.Errmsg)
	}
	return
}

func (s *Service) WithDrawTemplateMsg(req model.TemplateMsgWithDraw) error {
	query := url.Values{}
	query.Set("access_token", s.AccessToken())

	msg := model.TemplateMsg{
		Touser:     "oqeBd0fGbtYTmoVGhHzZ5Nf3-Egc",
		TemplateID: "TdBfxZjmJeshcJsUY6Y08hKLaSjn_A61MfcieezyQ8c",
		URL:        "",
		Data: model.WithDrawTemplateMsgData{
			First: model.TemplateMsgFirst{
				Value: req.OpenID,
				Color: "#4682b4",
			},
			OrderIDs: model.TemplateMsgKeyword1{
				Value: strings.Join(req.OrderIDs, "\n"),
				Color: "#4682b4",
			},
			NickName: model.TemplateMsgKeyword2{
				Value: req.NickName,
				Color: "#4682b4",
			},
			Rebate: model.TemplateMsgKeyword3{
				Value: req.Rebate,
				Color: "#4682b4",
			},
			WithDrawTime: model.TemplateMsgKeyword4{
				Value: req.WithDrawTime,
				Color: "#4682b4",
			},
			Action: model.TemplateMsgKeyword5{
				Value: req.Action,
				Color: "#4682b4",
			},
			Remark: model.TemplateMsgRemark{
				Value: "12期全免特权卡+善诊双人体验, 99元立抢>>",
				Color: "#FE4365",
			},
		},
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	v := new(struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Msgid   int    `json:"msgid"`
	})
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", s.AccessToken()), bytes.NewReader(marshal))
	if err != nil {
		return err
	}
	err = s.client.Do(context.Background(), request, v)
	if err != nil {
		return err
	}
	if v.Errcode != 0 {
		err = fmt.Errorf("错误代码: %d,错误信息: %s", v.Errcode, v.Errmsg)
	}
	return err
}

// #FE4365 红色   #4682b4 钢蓝色
func (s *Service) MatchedTemplateMsgSend(ctx context.Context, req *pb.MatchedTemplateMsgSendReq) (resp *empty.Empty, err error) {
	query := url.Values{}
	query.Set("access_token", s.AccessToken())
	msg := model.TemplateMsg{
		Touser:     req.UserID,
		TemplateID: "ZCrcPnmvT6xUMrILO-rVSM9EFg_s9QGEa-69BHPhK9c",
		URL:        "",
		Data: model.MatchedTemplateMsgData{
			First: model.TemplateMsgFirst{
				Value: "尊敬的用户: ",
				Color: "#4682b4",
			},
			Title: model.TemplateMsgKeyword1{
				Value: req.Title,
				Color: "#4682b4",
			},
			PaidTime: model.TemplateMsgKeyword2{
				Value: req.PaidTime,
				Color: "#4682b4",
			},
			OrderID: model.TemplateMsgKeyword3{
				Value: req.OrderID,
				Color: "#4682b4",
			},
			AlipayTotalPrice: model.TemplateMsgKeyword4{
				Value: req.AlipayTotalPrice,
				Color: "#4682b4",
			},
			Rebate: model.TemplateMsgKeyword5{
				Value: req.Rebate,
				Color: "#4682b4",
			},
			Remark: model.TemplateMsgRemark{
				Value: "12期全免特权卡+善诊双人体验, 99元立抢>>",
				Color: "#FE4365",
			},
		},
	}
	marshal, err := json.Marshal(msg)
	if err != nil {
		return
	}
	v := new(struct {
		Errcode int    `json:"errcode"`
		Errmsg  string `json:"errmsg"`
		Msgid   int    `json:"msgid"`
	})
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", s.AccessToken()), bytes.NewReader(marshal))
	if err != nil {
		return
	}
	err = s.client.Do(ctx, request, v)
	if err != nil {
		return
	}
	if v.Errcode != 0 {
		err = fmt.Errorf("错误代码: %d,错误信息: %s", v.Errcode, v.Errmsg)
	}
	return
}

func (s *Service) KeyConvert(ctx context.Context, req model.KeyConvertReq, args ...interface{}) (resp model.KeyConvertResp, err error) {
	option := make([]grpc.CallOption, len(args))
	for i := range args {
		if o, ok := args[i].(grpc.CallOption); ok {
			option[i] = o
		}
	}
	get := new(pb.KeyConvertResp)
	get, err = s.tbkClient.KeyConvert(ctx, &pb.KeyConvertReq{
		FromKey: req.FromKey,
		UserID:  req.UserID,
	}, option...)
	if err != nil {
		err = fmt.Errorf("KeyConvert: (%w)", err)
		return
	}
	if get == nil {
		err = errors.New("KeyConvert: rpc get is nil")
		return
	}
	resp.Price = get.Price
	resp.Rebate = get.Rebate
	resp.Coupon = get.Coupon
	resp.PicURL = get.PicURL
	resp.Title = get.Title
	resp.ItemURL = get.ItemURL
	return
}

func (s *Service) WithDraw(ctx context.Context, req model.WithDrawReq, args ...interface{}) (resp model.WithDrawResp, err error) {
	option := make([]grpc.CallOption, len(args))
	for i := range args {
		if o, ok := args[i].(grpc.CallOption); ok {
			option[i] = o
		}
	}
	get := new(pb.WithDrawResp)
	get, err = s.tbkClient.WithDraw(ctx, &pb.WithDrawReq{UserID: req.UserID}, option...)
	if err != nil {
		err = fmt.Errorf("WithDraw: (%w)", err)
		return
	}
	if get == nil {
		err = errors.New("WithDraw: rpc get is nil")
		return
	}
	resp.OrderIDs = get.OrderIDs
	resp.Rebate = get.Rebate
	return

}
func (s *Service) NewsURLGet(ctx context.Context, req model.NewsURLGetReq, args ...interface{}) (resp model.NewsURLGetResp, err error) {
	option := make([]grpc.CallOption, len(args))
	for i := range args {
		if o, ok := args[i].(grpc.CallOption); ok {
			option[i] = o
		}
	}
	get := new(pb.NewsURLGetResp)
	get, err = s.fileSystemClient.NewsURLGet(ctx, &pb.NewsURLGetReq{
		FakeID:    req.FakeID,
		Timestamp: req.Timestamp,
	}, option...)
	if err != nil {
		err = fmt.Errorf("NewsURLGet: (%w)", err)
		return
	}
	if get == nil {
		err = errors.New("NewsURLGet: rpc get is nil")
		return
	}
	resp.URL = get.URL
	return
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
		txMessage = s.handleText(ctx, text, rx.FromUserName, rx.CreateTime)
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
		return xml.Marshal(model.NewTxMessage(rx.FromUserName, rx.ToUserName, specifics.TxMessageType(), time.Now().Unix(), specifics))
	}

	tx := model.NewTxMessage(rx.FromUserName, rx.ToUserName, txMessage.TxMessageType(), time.Now().Unix(), txMessage)
	log.Info("reply message: %s", tx)
	return xml.Marshal(tx)

}
func (s *Service) handleText(ctx context.Context, text model.RxTextMessageExtra, openID string, createTime int64) model.TxMessageKind {
	deadline, _ := ctx.Deadline()
	log.Info("handleText 开始时过期时间: %d", time.Until(deadline))
	defer func() {
		deadline, _ := ctx.Deadline()
		log.Info("handleText 结束时过期时间: %d", time.Until(deadline))
	}()
	switch content := text.GetContent(); content {
	case "我要表情包":
		return model.NewTxTextMessageSpecifics("你可以尝试先向我发送表情包")
	case "我要消费券":
		return model.NewTxTextMessageSpecifics("你可以尝试先向我发送淘口令")
	case "提现":
		resp, err := s.WithDraw(ctx, model.WithDrawReq{UserID: openID})
		if err != nil {
			return model.NewTxTextMessageSpecifics(err.Error())
		}
		if err = s.WithDrawTemplateMsg(model.TemplateMsgWithDraw{
			OpenID:       openID,
			OrderIDs:     resp.OrderIDs,
			NickName:     "",
			Rebate:       resp.Rebate,
			WithDrawTime: time.Now().String(),
			Action:       "",
		}); err != nil {
			//回调
		}
		if len(resp.OrderIDs) == 0 {
			return model.NewTxTextMessageSpecifics("目前没有可提现订单,请查看订单是否确认收货.")
		}
		return model.NewTxTextMessageSpecifics("添加微信号: woquannibiemai, 微信昵称: c")
	case "【收到不支持的消息类型，暂无法显示】":
		newsURLGetResp, err := s.NewsURLGet(context.Background(), model.NewsURLGetReq{
			FakeID:    openID,
			Timestamp: createTime,
		})
		if err != nil {
			err = fmt.Errorf("ReplyMessage: (%w)", err)
		}
		return model.NewTxNewsMessageSpecifics(1, "测试啊啊啊啊", "啊实打实大苏打", "http://mmbiz.qpic.cn/sz_mmbiz_png/lPpqOpT1r9c7dqO5WvC2Py8zPAvBVJ3cCJwQMEKeNJ3eNRodPtkfic8C9b6o8AZwZWzeIkEuL6pQYSxmS459mHw/0?wx_fmt=png", newsURLGetResp.URL)
	default:
		submatch := tklRe.FindSubmatch([]byte(content))
		if len(submatch) > 1 {
			fromKey := string(submatch[1])

			keyConvertKeyResp, err := s.KeyConvert(ctx, model.KeyConvertReq{
				FromKey: fromKey,
				UserID:  openID,
			})
			if err != nil {
				log.Error("handleText failed: %+v", err)
				return model.NewTxTextMessageSpecifics(err.Error())
			}
			if keyConvertKeyResp.Coupon != "" {
				keyConvertKeyResp.Coupon = "优惠券:" + keyConvertKeyResp.Coupon
			}
			return model.NewTxNewsMessageSpecifics(1, fmt.Sprintf("约返:%s %s 付费价:%s", keyConvertKeyResp.Rebate, keyConvertKeyResp.Coupon, keyConvertKeyResp.Price), keyConvertKeyResp.Title, keyConvertKeyResp.PicURL, keyConvertKeyResp.ItemURL)
		}
		return model.NewTxTextMessageSpecifics(`<a href="weixin://bizmsgmenu?msgmenucontent=我要表情包&msgmenuid=1">我要表情包</a>			<a href="weixin://bizmsgmenu?msgmenucontent=我要消费券&msgmenuid=1">我要消费券</a>			<a href="weixin://bizmsgmenu?msgmenucontent=提现&msgmenuid=1">提现</a>`)
	}
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

func (s *Service) AccessToken() string {
	s.once.Do(func() {
		for s.accessToken == "" {
			timer := time.NewTimer(time.Second * 5)
			<-timer.C
			var err error
			if s.accessToken, err = s.execGetAccessToken(context.Background()); err != nil {
				log.Error("renovateToken error: (%v)", err)
			}
		}
		go s.renovateToken()

	})
	return s.accessToken
}
func (s *Service) renovateToken() {
	ctx := context.Background()
	for range time.Tick(time.Hour * 2) {
		var err error
		if s.accessToken, err = s.execGetAccessToken(ctx); err != nil {
			log.Error("renovateToken error: (%v)", err)
		}
	}
}

func (s *Service) execGetAccessToken(ctx context.Context) (accessToken string, err error) {
	appid := s.GetAppID()
	secret := s.GetSecret()
	baseURL := "https://api.weixin.qq.com/cgi-bin/token"
	query := url.Values{}
	query.Set("grant_type", "client_credential")
	query.Set("appid", appid)
	query.Set("secret", secret)
	request, err := s.client.NewRequest(http.MethodGet, baseURL, "", query)
	if err != nil {
		return
	}
	v := &model.AccessTokenResp{}
	err = s.client.JSON(ctx, request, v)
	if err != nil {
		return
	}
	if v.Code != 0 {
		err = fmt.Errorf("错误代码: %d,错误信息: %s", v.Code, v.ErrMsg)
		return
	}
	accessToken = v.AccessToken.AccessToken
	return
}

// New new a service and return.
func New(d dao.Dao, client *bm.Client) (s *Service, cf func(), err error) {
	s = &Service{
		ac:     &paladin.TOML{},
		dao:    d,
		client: client,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	if s.fileSystemClient, err = pb.NewfileSystemClient(&warden.ClientConfig{
		Timeout: xtime.Duration(time.Second * 5),
	}); err != nil {
		panic(err)
	}
	if s.tbkClient, err = pb.NewtbkClient(&warden.ClientConfig{
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
