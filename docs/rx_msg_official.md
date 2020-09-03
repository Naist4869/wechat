# 接收消息格式

## Models

### `RxMessage` 一条接收到的消息

  Name           | XML            | Type     | Doc
  :------------- | :------------- | :------- | :-----------------------
  `ToUserName`   | ``   | `string` | 开发者微信号
  `FromUserName` | `` | `string` | 发送方帐号（一个OpenID）
  `CreateTime`   | ``   | `int64`  | 消息创建时间 （整型）
  `MsgType`      | ``      | `RxMessageType` | 消息类型
  `MsgID`        | ``        | `int64`  | 消息id，64位整
  `extra`        | ``        | ` RxMessageKind` 里面放的是不同对象的实现类型 |

### `TxMessage` 一条要发送的消息

  Name           | XML            | Type     | Doc
  :------------- | :------------- | :------- | :-----------------------
  `XMLName`   | `xml`   | `xmlName` | XML头
  `ToUserName`   | `ToUserName`   | `cdataNode` | 开发者微信号
  `FromUserName` | `FromUserName` | `cdataNode` | 发送方帐号（一个OpenID）
  `CreateTime`   | `CreateTime`   | `int64`  | 消息创建时间 （整型）
  `MsgType`      | `MsgType`      | `cdataNode` | 消息类型，文本为text
  `Extra`        | ``        | ` TxMessageKind` 里面放的是不同对象的实现类型 |

### `rxMessageSuperset` 接收消息的超集

  Name           | XML            | Type     | Doc
  :------------- | :------------- | :------- | :-----------------------
  `ToUserName`   | `ToUserName`   | `string` | 开发者微信号
  `FromUserName` | `FromUserName` | `string` | 发送方帐号（一个OpenID）
  `CreateTime`   | `CreateTime`   | `int64`  | 消息创建时间 （整型）
  `MsgType`      | `MsgType`      | `RxMessageType` | 消息类型，文本为text
  `MsgID`        | `MsgId`        | `int64`  | 消息id，64位整
`Content`|`Content`|`string`|文本消息内容
`PicURL`|`PicUrl`|`string`|图片链接（由系统生成）
`MediaID`|`MediaId`|`string`|图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效,语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效,视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效,小视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`Format`|`Format`|`string`|语音格式，如amr，speex等
`Recognition`|`Recognition`|`string`|语音识别结果，UTF8编码
`ThumbMediaID`|`ThumbMediaId`|`string`|视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效,小视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内 
`Lat`|`Location_X`|`float64`|地理位置纬度
`Lon`|`Location_Y`|`float64`|地理位置经度
`Scale`|`Scale`|`int`|地图缩放大小
`Label`|`Label`|`string`|地理位置信息
`Title`|`Title`|`string`|标题
`Description`|`Description`|`string`|描述
`URL`|`Url`|`string`|链接跳转的url

### `rxMessageCommon` 接收消息的公共部分

  Name           | XML            | Type     | Doc
  :------------- | :------------- | :------- | :-----------------------
  `ToUserName`   | `ToUserName`   | `string` | 开发者微信号
  `FromUserName` | `FromUserName` | `string` | 发送方帐号（一个OpenID）
  `CreateTime`   | `CreateTime`   | `int64`  | 消息创建时间 （整型）
  `MsgType`      | `MsgType`      | `RxMessageType` | 消息类型，文本为text
  `MsgID`        | `MsgId`        | `int64`  | 消息id，64位整

### `txMessageCommon` 接收消息的公共部分

  Name           | XML            | Type     | Doc
  :------------- | :------------- | :------- | :-----------------------
  `XMLName`   | `xml`   | `xmlName` | XML头
  `ToUserName`   | `ToUserName`   | `cdataNode` | 开发者微信号
  `FromUserName` | `FromUserName` | `cdataNode` | 发送方帐号（一个OpenID）
  `CreateTime`   | `CreateTime`   | `int64`  | 消息创建时间 （整型）
  `MsgType`      | `MsgType`      | `cdataNode` | 消息类型，文本为text
  `MsgID`        | `MsgId`        | `int64`  | 消息id，64位整

 ```go
// RxMessageType 消息类型
type RxMessageType string

// RxMessageTypeText 文本消息
const RxMessageTypeText RxMessageType = "text"

// RxMessageTypeImage 图片消息
const RxMessageTypeImage RxMessageType = "image"

// RxMessageTypeVoice 语音消息
const RxMessageTypeVoice RxMessageType = "voice"

// RxMessageTypeVideo 视频消息
const RxMessageTypeVideo RxMessageType = "video"

// RxMessageTypeShortVideo 小视频消息
const RxMessageTypeShortVideo RxMessageType = "shortvideo"

// RxMessageTypeLocation 位置消息
const RxMessageTypeLocation RxMessageType = "location"

// RxMessageTypeLink 链接消息
const RxMessageTypeLink RxMessageType = "link"

// TxMessageType 消息类型
type TxMessageType string

// TxMessageTypeText 文本消息
const TxMessageTypeText TxMessageType = "text"

// TxMessageTypeImage 图片消息
const TxMessageTypeImage TxMessageType = "image"

// TxMessageTypeVoice 语音消息
const TxMessageTypeVoice TxMessageType = "voice"

// TxMessageTypeVideo 视频消息
const TxMessageTypeVideo TxMessageType = "video"

// TxMessageTypeMusic 音乐消息
const TxMessageTypeMusic TxMessageType = "music"

// TxMessageTypeNews 图文消息
const TxMessageTypeNews TxMessageType = "news"

// TxMessageTypeCS 转发客服消息
const TxMessageTypeCS TxMessageType = "transfer_customer_service"

const (
  ContentField = "Content"
  PicURLField= "PicUrl"
  MediaIDField="MediaId"
  FormatField = "Format"
  ThumbMediaIDField = "ThumbMediaId"
  LatField="Location_X"
  LonField = "Location_Y"
  ScaleField = "Scale"
  LabelField = "Label"
  TitleField = "Title"
  DescriptionField = "Description"
  URLField = "Url"
  RecognitionField="Recognition"
)
```

### `rxTextMessageSpecifics` 接收的文本消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Content`|`Content`|`string`|文本消息内容

### `txTextMessageSpecifics` 发送的文本消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `Content`   | `xmlName` | XML头
`Content`|`,cdata`|`string`|回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）

### `rxImageMessageSpecifics` 接收的图片消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`PicURL`|`PicUrl`|`string`|图片链接（由系统生成）
`MediaID`|`MediaId`|`string`|图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效

### `txImageMessageSpecifics` 发送的图片消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `Image`   | `xmlName` | XML头
`MediaID`|`MediaId`|`cdataNode`|通过素材管理中的接口上传多媒体文件，得到的id。

### `rxVoiceMessageSpecifics` 接收的语音消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`Format`|`Format`|`string`|语音格式，如amr，speex等
`Recognition`|`Recognition`|`string`|语音识别结果，UTF8编码


### `txVoiceMessageSpecifics` 发送的语音消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `Voice`   | `xmlName` | XML头
`MediaID`|`MediaId`|`cdataNode`|通过素材管理中的接口上传多媒体文件，得到的id



### `rxVideoMessageSpecifics` 接收的视频消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`ThumbMediaID`|`ThumbMediaId`|`string`|视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效

### `txVideoMessageSpecifics` 发送的视频消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `Video`   | `xmlName` | XML头
`MediaID`|`MediaId`|`cdataNode`|通过素材管理中的接口上传多媒体文件，得到的id
`Title`|`Title,omitempty`|`*cdataNode`|视频消息的标题
`Description`|`Description,omitempty`|`*cdataNode`|视频消息的描述


### `rxShortVideoMessageSpecifics` 接收的小视频消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`MediaID`|`MediaId`|`string`|小视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
`ThumbMediaID`|`ThumbMediaId`|`string`|小视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内

### `rxLocationMessageSpecifics` 接收的位置消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Lat`|`Location_X`|`float64`|地理位置纬度
`Lon`|`Location_Y`|`float64`|地理位置经度
`Scale`|`Scale`|`int`|地图缩放大小
`Label`|`Label`|`string`|地理位置信息

### `rxLinkMessageSpecifics` 接收的链接消息，特有字段

Name|XML|Type|Doc
:---|:--|:---|:--
`Title`|`Title`|`string`|标题
`Description`|`Description`|`string`|描述
`URL`|`Url`|`string`|链接跳转的url

### `txMusicMessageSpecifics` 发送的音乐消息，特有字段 MsgType 音乐为music

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `Music`   | `xmlName` | XML头
`Title`|`Title,omitempty`|`*cdataNode`|音乐标题
`Description`|`Description,omitempty`|`*cdataNode`|音乐描述
`MusicURL`|`MusicUrl,omitempty`|`*cdataNode`|音乐链接
`HQMusicURL`|`HQMusicUrl,omitempty`|`*cdataNode`|高质量音乐链接，WIFI环境优先使用该链接播放音乐
`ThumbMediaID`|`ThumbMediaId`|`cdataNode`|缩略图的媒体id，通过素材管理中的接口上传多媒体文件，得到的id,必须字段

### `txNewsMessageSpecifics` 发送的图文消息，特有字段 MsgType 图文为news

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `xml`   | `xmlName` | XML头
`ArticleCount`|`ArticleCount`|`int`|图文消息个数；当用户发送文本、图片、视频、图文、地理位置这五种消息时，开发者只能回复1条图文消息；其余场景最多可回复8条图文消息
`Articles`|`Articles`|`[]Items`|图文消息信息，注意，如果图文数超过限制，则将只发限制内的条数

### `Items` items

Name|XML|Type|Doc
:---|:--|:---|:--
`Item`   | `item` | `item`   | item


### `item` 图文消息类型
Name|XML|Type|Doc
:---|:--|:---|:--
`Title`|`Title`|`cdataNode`|图文消息标题
`Description`|`Description`|`cdataNode`|图文消息描述
`PicURL`|`PicUrl`|`cdataNode`|图片链接，支持JPG、PNG格式，较好的效果为大图360*200，小图200*200
`URL`|`Url`|`cdataNode`|点击图文消息跳转链接

### `txCSMessageSpecifics` 发送的客服消息，特有字段 MsgType 客服为transfer_customer_service

Name|XML|Type|Doc
:---|:--|:---|:--
`XMLName`   | `TransInfo`   | `xmlName` | XML头
`KfAccount`|`KfAccount`|`cdataNode`|指定会话接入的客服账号
