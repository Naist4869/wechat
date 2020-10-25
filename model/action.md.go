// Code generated by sdkcodegen; DO NOT EDIT.

package model

// TemplateMsg 模板消息
type TemplateMsg struct {
	// Touser 接收者 openid
	Touser string `json:"touser"`
	// TemplateID 模板 ID
	TemplateID string `json:"template_id"`
	// URL 模板跳转链接（海外帐号没有跳转能力）
	URL string `json:"url,omitempty"`
	// Miniprogram 跳小程序所需数据，不需跳小程序可不用传该数据
	Miniprogram TemplateMsgMiniprogram `json:"miniprogram,omitempty"`
	// Data 模板数据 {{first.DATA}}
	Data interface{} `json:"data"`
}

// TemplateMsgMiniprogram 模板消息
type TemplateMsgMiniprogram struct {
	// Appid 所需跳转到的小程序 appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）
	Appid string `json:"appid"`
	// Pagepath 所需跳转到小程序的具体页面路径，支持带参数,（示例 index?foo=bar），要求该小程序已发布，暂不支持小游戏
	Pagepath string `json:"pagepath,omitempty"`
}

// MatchedTemplateMsgData 跟单模板消息
type MatchedTemplateMsgData struct {
	// First {{first.DATA}}
	First TemplateMsgFirst `json:"first,omitempty"`
	// Title 商品名称
	Title TemplateMsgKeyword1 `json:"keyword1,omitempty"`
	// PaidTime 下单时间
	PaidTime TemplateMsgKeyword2 `json:"keyword2,omitempty"`
	// OrderID 订单编号
	OrderID TemplateMsgKeyword3 `json:"keyword3,omitempty"`
	// AlipayTotalPrice 订单金额
	AlipayTotalPrice TemplateMsgKeyword4 `json:"keyword4,omitempty"`
	// Rebate 预计返利
	Rebate TemplateMsgKeyword5 `json:"keyword5,omitempty"`
	// Remark {{remark.DATA}}
	Remark TemplateMsgRemark `json:"remark,omitempty"`
}

// BalanceTemplateMsgData 结算模板消息
type BalanceTemplateMsgData struct {
	// Title {{first.DATA}}
	Title TemplateMsgFirst `json:"first,omitempty"`
	// EarningTime 结算时间
	EarningTime TemplateMsgKeyword1 `json:"keyword1,omitempty"`
	// Salary 结算金额
	Salary TemplateMsgKeyword2 `json:"keyword2,omitempty"`
	// Balance 当前余额
	Balance TemplateMsgKeyword3 `json:"keyword3,omitempty"`
	// Remark {{remark.DATA}}
	Remark TemplateMsgRemark `json:"remark,omitempty"`
}

// WithDrawTemplateMsgData 提现模板消息
type WithDrawTemplateMsgData struct {
	// First {{first.DATA}}
	First TemplateMsgFirst `json:"first,omitempty"`
	// OrderIDs 订单号
	OrderIDs TemplateMsgKeyword1 `json:"keyword1,omitempty"`
	// NickName 昵称
	NickName TemplateMsgKeyword2 `json:"keyword2,omitempty"`
	// Rebate 金额
	Rebate TemplateMsgKeyword3 `json:"keyword3,omitempty"`
	// WithDrawTime 时间
	WithDrawTime TemplateMsgKeyword4 `json:"keyword4,omitempty"`
	// Action 方式
	Action TemplateMsgKeyword5 `json:"keyword5,omitempty"`
	// Remark {{remark.DATA}}
	Remark TemplateMsgRemark `json:"remark,omitempty"`
}

// TemplateMsgFirst 模板消息
type TemplateMsgFirst struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgKeyword1 模板消息
type TemplateMsgKeyword1 struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgKeyword2 模板消息
type TemplateMsgKeyword2 struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgKeyword3 模板消息
type TemplateMsgKeyword3 struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgKeyword4 模板消息
type TemplateMsgKeyword4 struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgRemark 模板消息
type TemplateMsgRemark struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}

// TemplateMsgKeyword5 模板消息
type TemplateMsgKeyword5 struct {
	// Value 内容
	Value string `json:"value"`
	// Color 模板内容字体颜色，不填默认为黑色
	Color string `json:"color,omitempty"`
}