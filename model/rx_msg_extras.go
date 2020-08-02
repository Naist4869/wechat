package model

import (
	"fmt"
	"io"
)

type RxMessageKind interface {
	messageKind
}

// NOTE: 这顺便就构成了一个封闭的 enum
type messageKind interface {
	formatInto(io.Writer)
}

var _ RxTextMessageExtra = (*rxTextMessageSpecifics)(nil)
var _ RxImageMessageExtra = (*rxImageMessageSpecifics)(nil)
var _ RxVoiceMessageExtra = (*rxVoiceMessageSpecifics)(nil)
var _ RxVideoMessageExtra = (*rxVideoMessageSpecifics)(nil)
var _ RxShortVideoMessageExtra = (*rxShortVideoMessageSpecifics)(nil)
var _ RxLocationMessageExtra = (*rxLocationMessageSpecifics)(nil)
var _ RxLinkMessageExtra = (*rxLinkMessageSpecifics)(nil)

// RxLinkMessageExtra 链接消息的参数。
type RxLinkMessageExtra interface {
	RxMessageKind

	// GetTitle 返回链接消息的标题。
	GetTitle() string

	// GetDescription 返回链接消息的描述。
	GetDescription() string

	// GetURL 返回链接消息的跳转 URL。
	GetURL() string
}

func (r *rxLinkMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"Title: %#v, Description: %#v, URL: %#v",
		r.Title,
		r.Description,
		r.URL,
	)
}

func (r *rxLinkMessageSpecifics) GetTitle() string {
	return r.Title
}

func (r *rxLinkMessageSpecifics) GetDescription() string {
	return r.Description
}

func (r *rxLinkMessageSpecifics) GetURL() string {
	return r.URL
}

// RxLocationMessageExtra 位置消息的参数。
type RxLocationMessageExtra interface {
	RxMessageKind

	// GetLatitude 返回位置消息的纬度（角度值；北纬为正）。
	GetLatitude() float64

	// GetLongitude 返回位置消息的经度（角度值；东经为正）。
	GetLongitude() float64

	// GetScale 返回位置消息的地图缩放大小。
	GetScale() int

	// GetLabel 返回位置消息的地理位置信息。
	GetLabel() string

	// 不知道这个有啥用，先不暴露
	// GetAppType() string
}

func (r *rxLocationMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"Latitude: %#v, Longitude: %#v, Scale: %d, Label: %#v",
		r.Lat,
		r.Lon,
		r.Scale,
		r.Label,
	)
}

func (r *rxLocationMessageSpecifics) GetLatitude() float64 {
	return r.Lat
}

func (r *rxLocationMessageSpecifics) GetLongitude() float64 {
	return r.Lon
}

func (r *rxLocationMessageSpecifics) GetScale() int {
	return r.Scale
}

func (r *rxLocationMessageSpecifics) GetLabel() string {
	return r.Label
}

type RxShortVideoMessageExtra RxVideoMessageExtra

func (r *rxShortVideoMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v, ThumbMediaID: %#v", r.MediaID, r.ThumbMediaID)
}

func (r *rxShortVideoMessageSpecifics) GetMediaID() string {
	return r.MediaID
}

func (r *rxShortVideoMessageSpecifics) GetThumbMediaID() string {
	return r.ThumbMediaID
}

// RxVideoMessageExtra 视频消息的参数。
type RxVideoMessageExtra interface {
	RxMessageKind

	// GetMediaID 返回视频消息的视频媒体文件 ID。
	//
	// 可以调用【获取媒体文件】接口拉取，仅三天内有效。
	GetMediaID() string

	// GetThumbMediaID 返回视频消息缩略图的媒体 ID。
	//
	// 可以调用【获取媒体文件】接口拉取，仅三天内有效。
	GetThumbMediaID() string
}

func (r *rxVideoMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v, ThumbMediaID: %#v", r.MediaID, r.ThumbMediaID)
}

func (r *rxVideoMessageSpecifics) GetMediaID() string {
	return r.MediaID
}

func (r *rxVideoMessageSpecifics) GetThumbMediaID() string {
	return r.ThumbMediaID
}

// RxVoiceMessageExtra 语音消息的参数。
type RxVoiceMessageExtra interface {
	RxMessageKind

	// GetMediaID 返回语音消息的语音媒体文件 ID。
	//
	// 可以调用【获取媒体文件】接口拉取，仅三天内有效。
	GetMediaID() string

	// GetFormat 返回语音消息的语音格式，如 "amr"、"speex" 等。
	GetFormat() string

	// GetRecognition 语音识别结果，UTF8编码  默认不开启
	GetRecognition() string
}

func (r *rxVoiceMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v, Format: %#v", r.MediaID, r.Format)
}

func (r *rxVoiceMessageSpecifics) GetMediaID() string {
	return r.MediaID
}

func (r *rxVoiceMessageSpecifics) GetFormat() string {
	return r.Format
}

func (r *rxVoiceMessageSpecifics) GetRecognition() string {
	return r.Recognition
}

// RxImageMessageExtra 图片消息的参数。
type RxImageMessageExtra interface {
	RxMessageKind

	// GetPicURL 返回图片消息的图片链接 URL。
	GetPicURL() string

	// GetMediaID 返回图片消息的图片媒体文件 ID。
	//
	// 可以调用【获取媒体文件】接口拉取，仅三天内有效。
	GetMediaID() string
}

func (r *rxImageMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "PicURL: %#v, MediaID: %#v", r.PicURL, r.MediaID)
}

func (r *rxImageMessageSpecifics) GetPicURL() string {
	return r.PicURL
}

func (r *rxImageMessageSpecifics) GetMediaID() string {
	return r.MediaID
}

// RxTextMessageExtra 文本消息的参数。
type RxTextMessageExtra interface {
	RxMessageKind

	// GetContent 返回文本消息的内容。
	GetContent() string
}

func (r *rxTextMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "Content: %#v", r.Content)
}

func (r *rxTextMessageSpecifics) GetContent() string {
	return r.Content
}
