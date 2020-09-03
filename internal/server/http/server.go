package http

import (
	"crypto/subtle"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
	"wechat/internal/bytesconv"
	"wechat/internal/service"

	"github.com/go-kratos/kratos/pkg/net/http/blademaster/binding"

	pb "wechat/api"
	"wechat/model"

	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
	bm "github.com/go-kratos/kratos/pkg/net/http/blademaster"
)

var svr service.Server

type handleXMLMsg func(ctx *bm.Context, replyMessage []byte, timestamp time.Time)

// New new a bm server.
func New(s service.Server) (engine *bm.Engine, err error) {
	var (
		cfg struct {
			Server *bm.ServerConfig
			Client *bm.ClientConfig
		}
		ct paladin.TOML
	)

	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("bm").UnmarshalTOML(&cfg); err != nil {
		return
	}
	svr = s
	engine = bm.DefaultServer(cfg.Server)
	pb.RegisterDemoBMServer(engine, s)
	initRouter(engine)
	err = engine.Start()
	return
}

type Template struct {
	// Title 连衣裙	商品标题
	Title string
	// PictURL http://gi4.md.alicdn.com/bao/uploaded/i4/xxx.jpg	商品主图
	PictURL string
	// ReservePrice 102.00	商品一口价格
	ReservePrice string
	// ZkFinalPrice 	88.00	折扣价（元） 若属于预售商品，付定金时间内，折扣价=预售价
	ZkFinalPrice string
	// UserType 1	卖家类型，0表示集市，1表示商城
	UserType int
	// Nick xx旗舰店	店铺名称
	Nick string
	// 高佣淘口令
	TKL string
}

func initRouter(e *bm.Engine) {
	e.Ping(ping)
	g := e.Group("/wechat")
	{
		g.GET("/start", howToStart)
		g.GET("/callback", certification)
		g.POST("/callback", verify, callback)
	}
}

func ping(ctx *bm.Context) {
	if _, err := svr.Ping(ctx, nil); err != nil {
		log.Error("ping error(%v)", err)
		ctx.AbortWithStatus(http.StatusServiceUnavailable)
	}
}

func certification(ctx *bm.Context) {
	v := new(struct {
		Signature string `form:"signature" validate:"required"`
		TimeStamp string `form:"timestamp" validate:"required"`
		Nonce     string `form:"nonce" validate:"required"`
		EchoStr   string `form:"echostr" validate:"required"`
	})
	if err := ctx.Bind(v); err != nil {
		return
	}
	token := svr.GetToken()

	if token == "" {
		err := errors.New("token was not set for server. ")
		log.Error("Certification: (%v)", err)
		return
	}
	wantSignature := Sign(token, v.TimeStamp, v.Nonce)
	if subtle.ConstantTimeCompare([]byte(v.Signature), []byte(wantSignature)) != 1 {
		log.Error("Certification check signature error! have: %s, want: %s", v.Signature, wantSignature)
		return
	}
	ctx.String(http.StatusOK, "%s", v.EchoStr)
}

func verify(ctx *bm.Context) {
	v := new(struct {
		Signature    string `form:"signature" binding:"required"`
		TimeStamp    string `form:"timestamp" binding:"required"`
		Nonce        string `form:"nonce" binding:"required"`
		OpenID       string `form:"openid" binding:"required"`
		EncryptType  string `form:"encrypt_type"`
		MsgSignature string `form:"msg_signature"`
	})
	if err := ctx.BindWith(v, binding.Query); err != nil {
		return
	}
	token := svr.GetToken()
	if token == "" {
		err := errors.New("token was not set for server. ")
		log.Error("Verify: (%v)", err)
		return
	}
	body, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Error("Verify ioutil.ReadAll error(%+v)", err)
		return
	}
	ctx.Request.Body.Close()
	RespErr := func(err error) {
		log.Error("Verify: (%v)", err)
		ctx.String(http.StatusOK, "success")
	}
	// 解密
	switch v.EncryptType {
	case "aes":
		xmlRxEncrypt := new(struct {
			ToUserName string `xml:"ToUserName"`
			Encrypt    string `xml:"Encrypt"`
		})

		if err := xml.Unmarshal(body, xmlRxEncrypt); err != nil {
			RespErr(fmt.Errorf("xmlRxEncrypt unmarshal fail(%w)", err))
			return
		}
		{
			haveToUserName := xmlRxEncrypt.ToUserName
			wantToUserName := svr.GetOriID()
			if strings.Compare(haveToUserName, wantToUserName) != 0 {
				RespErr(fmt.Errorf("the message ToUserName mismatch, have: %s, want: %s", haveToUserName, wantToUserName))
				return
			}
		}
		{
			wantMsgSignature := bytesconv.StringToBytes(Sign(token, v.TimeStamp, v.Nonce, xmlRxEncrypt.Encrypt))
			if subtle.ConstantTimeCompare(bytesconv.StringToBytes(v.MsgSignature), wantMsgSignature) != 1 {
				RespErr(fmt.Errorf("check msg_signature failed! have: %s, want: %s", v.MsgSignature, wantMsgSignature))
				return
			}
		}

		aesKey := svr.GetAESKey()
		random, rawXMLMsg, haveAppIdBytes, err := AESDecryptMsg(bytesconv.StringToBytes(xmlRxEncrypt.Encrypt), aesKey)
		if err != nil {
			RespErr(err)
			return
		}
		wantAppId := svr.GetAppID()
		haveAppId := string(haveAppIdBytes)
		if len(wantAppId) != 0 && strings.Compare(haveAppId, wantAppId) != 0 {
			err := fmt.Errorf("the message AppId mismatch, have: %s, want: %s", haveAppId, wantAppId)
			RespErr(err)
			return
		}
		ctx.Set("rawXMLMsg", rawXMLMsg)
		ctx.Set("handleXMLMsg", handleXMLMsg(func(ctx *bm.Context, replyMessage []byte, timestamp time.Time) {
			encrypt, err := AESEncryptMsg(random, replyMessage, wantAppId, aesKey)
			if err != nil {
				log.Error("handleXMLMsg (%v)", err)
				return
			}
			unix := timestamp.Unix()
			nonce := makeNonce()
			timestampStr := strconv.FormatInt(unix, 10)
			msgSignature := Sign(token, timestampStr, nonce, encrypt)
			ctx.XML(model.XMLTxEncryptEnvelope{
				Encrypt:      encrypt,
				MsgSignature: msgSignature,
				Timestamp:    unix,
				Nonce:        nonce,
			}, nil)
		}))
	case "raw", "":
		ctx.Set("rawXMLMsg", body)
		ctx.Set("handleXMLMsg", handleXMLMsg(func(ctx *bm.Context, replyMessage []byte, timestamp time.Time) {
			ctx.Bytes(http.StatusOK, "application/xml; charset=utf-8", replyMessage)
		}))
	default:
		RespErr(errors.New("unknown encrypt_type: " + v.EncryptType))
		return
	}
}
func callback(ctx *bm.Context) {
	rawXMLMsg, exists := ctx.Get("rawXMLMsg")
	if !exists {
		ctx.JSON(nil, errors.New("UseCase: missing required parameters"))
		return
	}
	handleFun, exists := ctx.Get("handleXMLMsg")
	if !exists {
		ctx.JSON(nil, errors.New("UseCase: missing required parameters"))
		return
	}
	log.Warn("receive raw message: \n%s", rawXMLMsg)
	replyMessage, err := svr.ReplyMessage(ctx, rawXMLMsg.([]byte))
	if err != nil {
		log.Error("callback: (%v)", err)
	}
	log.Warn("reply raw message: %s", replyMessage)
	handleFun.(handleXMLMsg)(ctx, replyMessage, time.Now())
}

// example for http request handler.
func howToStart(c *bm.Context) {
	k := &model.Kratos{
		Hello: "Golang 大法好 !!!",
	}
	c.JSON(k, nil)
}
