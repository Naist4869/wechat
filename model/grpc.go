package model

type MediaIDReq struct {
	FakeID    string
	Timestamp int64
}

type MediaIDResp struct {
	MediaID string
}
type NewsURLGetReq struct {
	FakeID    string
	Timestamp int64
}
type NewsURLGetResp struct {
	URL string
}
type KeyConvertReq struct {
	FromKey string
	UserID  string
}
type KeyConvertResp struct {
	Price   string
	Rebate  string
	Coupon  string
	Title   string
	PicURL  string
	ItemURL string
}

type WithDrawReq struct {
	UserID string
}

type WithDrawResp struct {
	Rebate   string
	OrderIDs []string
}
