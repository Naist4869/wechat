package model

import "encoding/xml"

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
