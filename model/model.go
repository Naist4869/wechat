package model

import (
	"encoding/xml"
	"time"
)

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

type XMLTxEncryptEnvelope struct {
	XMLName      xml.Name `xml:"xml"`
	Encrypt      string   `xml:"Encrypt"`
	MsgSignature string   `xml:"MsgSignature"`
	Timestamp    int64    `xml:"TimeStamp"`
	Nonce        string   `xml:"Nonce"`
}
type respCommon struct {
	Code   int    `json:"errcode"`
	ErrMsg string `json:"errMsg"`
}

type AccessTokenResp struct {
	respCommon
	AccessToken
}

type AccessToken struct {
	AccessToken   string        `json:"access_token"`
	ExpiresInSecs time.Duration `json:"expires_in"`
}
type TemplateMsgWithDraw struct {
	OpenID       string   // 申请用户的ID
	OrderIDs     []string // 订单号
	NickName     string   // 昵称
	Rebate       string   // 金额
	WithDrawTime string   //时间
	Action       string   // 方式
}
