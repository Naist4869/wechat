package model

import (
	"encoding/xml"
)

type cdataNode struct {
	CData string `xml:",cdata"`
}
type xmlName = xml.Name

func NewTxMessage(toUserName, fromUserName string, msgType TxMessageType, createTime int64, msgID int64, extra TxMessageKind) *TxMessage {
	return &TxMessage{
		ToUserName:   cdataNode{toUserName},
		FromUserName: cdataNode{fromUserName},
		CreateTime:   createTime,
		MsgType:      cdataNode{string(msgType)},
		MsgID:        msgID,
		Extra:        extra,
	}
}
