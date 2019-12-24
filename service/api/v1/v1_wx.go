package v1

import (
	"errors"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"go-memo/pkg/wx/wx_client"
	sdk "go-memo/pkg/wx/wx_sdk"
	"go-memo/service/api/viewset"
)

type WxViewset struct {
	viewset.ViewSet
}

func (this *WxViewset) ErrorHandler(f func(c *gin.Context) error) func(c *gin.Context) {
	return func(c *gin.Context) {
		this.ViewSet.ErrorHandler(f, c)
	}
}


// @Summary weixin check
// @Description 微信服务器校验
// @Accept  json
// @Produce  json
// @Param  signature  timestamp  echostr
// @Success 200 {object} viewset.Response
// @Router /api/public/wx/ [get]
func (this *WxViewset) Weixin(c *gin.Context) (err error) {
	s := sdk.TokenSignature{}
	s.Signature = c.Query("signature")
	s.Timestamp = com.StrTo(c.Query("timestamp")).MustInt()
	s.Echostr = c.Query("echostr")
	s.Nonce = com.StrTo(c.Query("nonce")).MustInt()

	if s.Confirm() {
		return this.SuccessResponse(c, s.Echostr)
	} else {
		return errors.New("wrong echostr")
	}
}

func (this *WxViewset) Login(c *gin.Context) (err error) {
	lp := &wx_client.LoginParams{}
	err = c.ShouldBind(lp)
	if err != nil {
		return
	}
	wx := wx_client.GetWxClient()
	res, err := wx.Login(lp)
	if err != nil {
		return
	}
	return this.SuccessResponse(c, res)
}

func (this *WxViewset) GetUserInfo(c *gin.Context) (err error) {
	return
}