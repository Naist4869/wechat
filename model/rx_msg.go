package model

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

func FromEnvelope(body []byte) (rxMessage *RxMessage, err error) {
	v := new(rxMessageSuperset)
	if err = xml.Unmarshal(body, v); err != nil {
		err = fmt.Errorf("FromEnvelope: (%w)", err)
		return
	}
	rxMessage = &RxMessage{
		FromUserName: v.FromUserName,
		ToUserName:   v.ToUserName,
		CreateTime:   v.CreateTime,
		MsgType:      v.MsgType,
		MsgID:        v.MsgID,
	}
	switch v.MsgType {
	case RxMessageTypeText:
		rxMessage.extra = &rxTextMessageSpecifics{Content: v.Content}
	case RxMessageTypeImage:
		rxMessage.extra = &rxImageMessageSpecifics{
			PicURL:  v.PicURL,
			MediaID: v.MediaID,
		}
	case RxMessageTypeVoice:
		rxMessage.extra = &rxVoiceMessageSpecifics{
			MediaID:     v.MediaID,
			Format:      v.Format,
			Recognition: v.Recognition,
		}
	case RxMessageTypeVideo:
		rxMessage.extra = &rxVideoMessageSpecifics{
			MediaID:      v.MediaID,
			ThumbMediaID: v.ThumbMediaID,
		}
	case RxMessageTypeLocation:
		rxMessage.extra = &rxLocationMessageSpecifics{
			Lat:   v.Lat,
			Lon:   v.Lon,
			Scale: v.Scale,
			Label: v.Label,
		}
	case RxMessageTypeLink:
		rxMessage.extra = &rxLinkMessageSpecifics{
			Title:       v.Title,
			Description: v.Description,
			URL:         v.URL,
		}
	case RxMessageTypeShortVideo:
		rxMessage.extra = &rxShortVideoMessageSpecifics{
			MediaID:      v.MediaID,
			ThumbMediaID: v.ThumbMediaID,
		}
	default:
		return nil, fmt.Errorf("FromEnvelope: unknown message type '%s'", v.MsgType)
	}
	return
}

// Text 如果消息为文本类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Text() (RxTextMessageExtra, bool) {
	y, ok := rx.extra.(RxTextMessageExtra)
	return y, ok
}

// Image 如果消息为图片类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Image() (RxImageMessageExtra, bool) {
	y, ok := rx.extra.(RxImageMessageExtra)
	return y, ok
}

// Voice 如果消息为语音类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Voice() (RxVoiceMessageExtra, bool) {
	y, ok := rx.extra.(RxVoiceMessageExtra)
	return y, ok
}

// Video 如果消息为视频类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Video() (RxVideoMessageExtra, bool) {
	y, ok := rx.extra.(RxVideoMessageExtra)
	return y, ok
}

// Location 如果消息为位置类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Location() (RxLocationMessageExtra, bool) {
	y, ok := rx.extra.(RxLocationMessageExtra)
	return y, ok
}

// Link 如果消息为链接类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) Link() (RxLinkMessageExtra, bool) {
	y, ok := rx.extra.(RxLinkMessageExtra)
	return y, ok
}

// ShortVideo 如果消息为小视频类型，则拿出相应的消息参数，否则返回 nil, false
func (rx RxMessage) ShortVideo() (RxShortVideoMessageExtra, bool) {
	y, ok := rx.extra.(RxShortVideoMessageExtra)
	return y, ok
}

func (tx *TxMessage) SetExtra(extra TxMessageKind) {
	tx.Extra = extra
}
func (rx *RxMessage) String() string {
	var sb strings.Builder

	_, _ = fmt.Fprintf(
		&sb,
		"RxMessage { FromUserName: %#v, ToUserName: %#v, CreateTime: %s, MsgType: %#v, MsgID: %d ",
		rx.FromUserName,
		rx.ToUserName,
		time.Unix(rx.CreateTime, 0).Format(time.RFC3339),
		rx.MsgType,
		rx.MsgID,
	)

	rx.extra.formatInto(&sb)

	sb.WriteString(" }")

	return sb.String()
}
func (tx *TxMessage) String() string {
	var sb strings.Builder

	_, _ = fmt.Fprintf(
		&sb,
		"TxMessage { FromUserName: %#v, ToUserName: %#v, CreateTime: %s, MsgType: %#v, MsgID: %d ",
		tx.FromUserName,
		tx.ToUserName,
		time.Unix(tx.CreateTime, 0).Format(time.RFC3339),
		tx.MsgType,
		tx.MsgID,
	)

	tx.Extra.formatInto(&sb)

	sb.WriteString(" }")

	return sb.String()
}
