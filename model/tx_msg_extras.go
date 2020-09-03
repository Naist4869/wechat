package model

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
)

type TxMessageKind interface {
	messageKind
	TxMessageType() TxMessageType
}

var _ TxMessageKind = (*txMusicMessageSpecifics)(nil)
var _ TxMessageKind = (*txCSMessageSpecifics)(nil)
var _ TxMessageKind = (*txNewsMessageSpecifics)(nil)
var _ TxMessageKind = (*txVideoMessageSpecifics)(nil)
var _ TxMessageKind = (*txVoiceMessageSpecifics)(nil)
var _ TxMessageKind = (*txTextMessageSpecifics)(nil)
var _ TxMessageKind = (*txImageMessageSpecifics)(nil)

func NewTxImageMessageSpecifics(mediaID string) *txImageMessageSpecifics {
	return &txImageMessageSpecifics{
		MediaID: cdataNode{CData: mediaID},
	}
}

func (t *txImageMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v", t.MediaID.CData)
}

func (t *txImageMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeImage
}

func NewTxTextMessageSpecifics(content string) TxMessageKind {
	return &txTextMessageSpecifics{
		Content: content,
	}
}

func (t *txTextMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeText
}

func (t *txTextMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "Content: %#v", t.Content)
}

func NewTxVoiceMessageSpecifics(mediaID string) *txVoiceMessageSpecifics {
	return &txVoiceMessageSpecifics{MediaID: cdataNode{mediaID}}
}
func (t *txVoiceMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v", t.MediaID.CData)
}

func (t *txVoiceMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeVoice
}

func (t *txNewsMessageSpecifics) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if t == nil {
		return errors.New("空指针异常")
	}
	if err := e.EncodeElement(t.ArticleCount, xml.StartElement{
		Name: xml.Name{
			Local: "ArticleCount",
		},
	}); err != nil {
		return err
	}

	return e.EncodeElement(t.Articles, xml.StartElement{
		Name: xml.Name{
			Local: "Articles",
		},
	})
}

func NewTxNewsMessageSpecifics(articleCount int, title string, description string, picURL string, url string) *txNewsMessageSpecifics {
	return &txNewsMessageSpecifics{
		ArticleCount: articleCount,
		Articles: []Items{{Item: item{
			Title:       cdataNode{title},
			Description: cdataNode{description},
			PicURL:      cdataNode{picURL},
			URL:         cdataNode{url},
		}}},
	}
}
func (t *txNewsMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"ArticleCount: %#v",
		t.ArticleCount,
	)
	for i := 0; i < t.ArticleCount; i++ {
		_, _ = fmt.Fprintf(
			w,
			"item { Title: %#v, Description: %#v, PicURL: %#v, URL: %#v",
			t.Articles[i].Item.Title.CData,
			t.Articles[i].Item.Description.CData,
			t.Articles[i].Item.PicURL.CData,
			t.Articles[i].Item.URL.CData,
		)
		_, _ = fmt.Fprint(
			w,
			" }")
	}

}

func (t *txNewsMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeNews
}

func NewTxVideoMessageSpecifics(mediaID, title, description string) *txVideoMessageSpecifics {
	t := &txVideoMessageSpecifics{
		MediaID: cdataNode{mediaID},
	}
	if title != "" {
		t.Title = &cdataNode{title}
	}
	if description != "" {
		t.Description = &cdataNode{description}
	}

	return t
}
func (t *txVideoMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(w, "MediaID: %#v, Title: %#v, Description: %#v", t.MediaID.CData, t.Title.CData, t.Description.CData)
}

func (t *txVideoMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeVideo
}

func NewTxMusicMessageSpecifics(thumbMediaID string, title, description, musicURL, HQMusicURL string) *txMusicMessageSpecifics {
	t := &txMusicMessageSpecifics{
		ThumbMediaID: cdataNode{thumbMediaID},
	}
	if title != "" {
		t.Title = &cdataNode{title}
	}
	if description != "" {
		t.Description = &cdataNode{description}
	}
	if musicURL != "" {
		t.MusicURL = &cdataNode{musicURL}
	}
	if HQMusicURL != "" {
		t.HQMusicURL = &cdataNode{HQMusicURL}
	}

	return t
}
func (t *txMusicMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"Title: %#v, Description: %#v, MusicURL: %#v, HQMusicURL: %#v, ThumbMediaID: %#v",
		t.Title.CData,
		t.Description.CData,
		t.MusicURL.CData,
		t.HQMusicURL.CData,
		t.ThumbMediaID.CData,
	)
}

func (t *txMusicMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeMusic
}

func NewTxCSMessageSpecifics(kfAccount string) *txCSMessageSpecifics {
	return &txCSMessageSpecifics{KfAccount: cdataNode{kfAccount}}
}

func (t *txCSMessageSpecifics) formatInto(w io.Writer) {
	_, _ = fmt.Fprintf(
		w,
		"KfAccount: %#v",
		t.KfAccount.CData,
	)
}

func (t *txCSMessageSpecifics) TxMessageType() TxMessageType {
	return TxMessageTypeCS
}
