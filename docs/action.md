# 微信公众号 api 需要的参数

## Models

### `TemplateMsg` 模板消息

| Name          | Type                     | JSON                    | Doc                                          |
| :------------ | :----------------------- | :---------------------- | :------------------------------------------- |
| `Touser`      | `string`                 | `touser`                | 接收者 openid                                |
| `TemplateID`  | `string`                 | `template_id`           | 模板 ID                                      |
| `URL`         | `string`                 | `url,omitempty`         | 模板跳转链接（海外帐号没有跳转能力）         |
| `Miniprogram` | `TemplateMsgMiniprogram` | `miniprogram,omitempty` | 跳小程序所需数据，不需跳小程序可不用传该数据 |
| `Data`        | `interface{}`        | `data`                  | 模板数据 {{first.DATA}}                      |

订单号：{{keyword1.DATA}}
支付时间：{{keyword2.DATA}}
支付金额：{{keyword3.DATA}}
商品名称：{{keyword4.DATA}}
{{remark.DATA}}

### `TemplateMsgMiniprogram` 模板消息

| Name       | Type     | JSON                 | Doc                                                                                                   |
| :--------- | :------- | :------------------- | :---------------------------------------------------------------------------------------------------- |
| `Appid`    | `string` | `appid`              | 所需跳转到的小程序 appid（该小程序 appid 必须与发模板消息的公众号是绑定关联关系，暂不支持小游戏）     |
| `Pagepath` | `string` | `pagepath,omitempty` | 所需跳转到小程序的具体页面路径，支持带参数,（示例 index?foo=bar），要求该小程序已发布，暂不支持小游戏 |

### `MatchedTemplateMsgData` 跟单模板消息

| Name       | Type                  | JSON                 | Doc               |
| :--------- | :-------------------- | :------------------- | :---------------- |
| `First`    | `TemplateMsgFirst`    | `first,omitempty`              | {{first.DATA}}    |
| `Title` | `TemplateMsgKeyword1` | `keyword1,omitempty` | 商品名称 |
| `PaidTime` | `TemplateMsgKeyword2` | `keyword2,omitempty` | 下单时间 |
| `OrderID` | `TemplateMsgKeyword3` | `keyword3,omitempty` | 订单编号 |
| `AlipayTotalPrice` | `TemplateMsgKeyword4` | `keyword4,omitempty` | 订单金额 |
| `Rebate` | `TemplateMsgKeyword5` | `keyword5,omitempty` | 预计返利 |
| `Remark`   | `TemplateMsgRemark`   | `remark,omitempty`             | {{remark.DATA}}   |

### `BalanceTemplateMsgData` 结算模板消息

| Name       | Type                  | JSON                 | Doc               |
| :--------- | :-------------------- | :------------------- | :---------------- |
| `Title`    | `TemplateMsgFirst`    | `first,omitempty`              | {{first.DATA}}    |
| `EarningTime` | `TemplateMsgKeyword1` | `keyword1,omitempty` | 结算时间 |
| `Salary` | `TemplateMsgKeyword2` | `keyword2,omitempty` | 结算金额 |
| `Balance` | `TemplateMsgKeyword3` | `keyword3,omitempty` | 当前余额 |
| `Remark`   | `TemplateMsgRemark`   | `remark,omitempty`             | {{remark.DATA}}   |

### `WithDrawTemplateMsgData` 提现模板消息

| Name       | Type                  | JSON                 | Doc               |
| :--------- | :-------------------- | :------------------- | :---------------- |
| `First`    | `TemplateMsgFirst`    | `first,omitempty`              | {{first.DATA}}    |
| `OrderIDs` | `TemplateMsgKeyword1` | `keyword1,omitempty` | 订单号 |
| `NickName` | `TemplateMsgKeyword2` | `keyword2,omitempty` | 昵称 |
| `Rebate` | `TemplateMsgKeyword3` | `keyword3,omitempty` | 金额 |
| `WithDrawTime` | `TemplateMsgKeyword4` | `keyword4,omitempty` | 时间 |
| `Action` | `TemplateMsgKeyword5` | `keyword5,omitempty` | 方式 |
| `Remark`   | `TemplateMsgRemark`   | `remark,omitempty`             | {{remark.DATA}}   |

### `TemplateMsgFirst` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgKeyword1` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgKeyword2` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgKeyword3` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgKeyword4` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgRemark` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |

### `TemplateMsgKeyword5` 模板消息

| Name    | Type     | JSON              | Doc                              |
| :------ | :------- | :---------------- | :------------------------------- |
| `Value` | `string` | `value`           | 内容                             |
| `Color` | `string` | `color,omitempty` | 模板内容字体颜色，不填默认为黑色 |