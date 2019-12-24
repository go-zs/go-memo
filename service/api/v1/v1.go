package v1

import (
	"go-memo/db/query"
	"go-memo/service/api/viewset"
)

func GetAuthVS() *AuthViewset {
	vs := &viewset.ViewSet{}
	authVS := &AuthViewset{
		itemInter: query.UserQ,
		ViewSet:   *vs,
	}
	return authVS
}

func GetWxVs() *WxViewset {
	vs := &viewset.ViewSet{}
	wxVs := &WxViewset{
		ViewSet: *vs,
	}
	return wxVs
}
