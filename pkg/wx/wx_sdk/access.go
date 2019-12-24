package sdk

import (
	"go-memo/conf"
	"go-memo/util"
	"sort"
	"strconv"
	"strings"
)



type TokenSignature struct {
	Signature string `json:"signature"`
	Timestamp int `json:"timestamp"`
	Nonce int `json:"nonce"`
	Echostr string `json:"echostr"`
}

// 签名验证
func (ts *TokenSignature) Confirm() bool {
	wxConfig := conf.GetConfig().Weixin
	token := wxConfig.Token
	list := []string{strconv.Itoa(ts.Nonce), strconv.Itoa(ts.Timestamp), token}
	sort.Strings(list)
	confirmStr := strings.Join(list, "")
	shaHash := util.Sha1(confirmStr)
	if shaHash == ts.Signature {
		return true
	}
	return false
}



