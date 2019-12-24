package sdk

import (
	"encoding/xml"
	"time"
)

const (
	text = "text"
	image = "image"
	voice = "voice"
	video = "video"
)

type Message struct {
	ToUserName   string // 开发者微信号
	FromUserName string // 发送方帐号（一个OpenID）
	CreateTime   int64  // 消息创建时间 （整型）
	MsgType      string // 消息类型，图片为image
	MsgId        int64  // 消息id，64位整型
	// text
	Content string // 文本消息内容
	// image
	PicUrl string // 图片链接（由系统生成）
	// voice
	Format string // 语音格式，如amr，speex等
	// video
	ThumbMediaId int64 // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。

	MediaId int64 // 图片消息媒体id，可以调用获取临时素材接口拉取数据。
}

type Reply struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string // 接收方帐号（收到的OpenID）
	FromUserName string //  开发者微信号
	CreateTime   int64  // 消息创建时间 （整型）
	MsgType      string // 消息类型，图片为image
	Image        Image
	Content      string // 文本消息内容  text
}

type Image struct {
	MediaId int64 // 通过素材管理中的接口上传多媒体文件，得到的id。
}

func Parse(s []byte) (m Message) {
	err := xml.Unmarshal(s, &m)
	if err != nil {
		return
	} else {
		return
	}
}

func (this *Message) Reply() (r Reply) {
	r = Reply{
		ToUserName:   this.FromUserName,
		FromUserName: this.ToUserName,
		CreateTime:   time.Now().Unix(),
	}

	switch this.MsgType {
	case text:
		this.ReplyText(&r)
	case image:
		this.ReplyImage(&r)
	default:
		r.MsgType = text
		r.Content = "你可以发送关键字给我"
	}
	return r

}

func (this *Message) ReplyText(r *Reply) {
	r.Content = "系统升级中"
	r.MsgType = text
	return
}

func (this *Message) ReplyImage(r *Reply) {
	r.Content = "系统升级中"
	r.MsgType = image
	return
}