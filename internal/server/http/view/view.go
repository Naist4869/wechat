package view

//Response 请求应答
type Response struct {
	Data     interface{} `json:"data"`     //数据
	Error    int         `json:"error"`    //错误码
	ErrorMsg string      `json:"errorMsg"` //错误信息
}
