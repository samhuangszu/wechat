package promotion

import (
	"github.com/samhuangszu/wechat/mch/core"
)

//Transfers 企业付款.
//  NOTE: 请求需要双向证书
func Transfers(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/promotion/transfers", req)
}

// Req 付款到零钱
type Req struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	MchAppID       string   `xml:"mch_appid" json:"mch_appid"`
	MchID          string   `xml:"mch_id" json:"mch_id"`
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	Sign           string   `xml:"sign" json:"sign"`
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	OpenID         string   `xml:"openid" json:"openid"`
	CheckName      string   `xml:"check_name" json:"check_name"`
	Amount         int      `xml:"amount" json:"amount,string"`
	Desc           string   `xml:"desc" json:"desc"`
	SpbillCreateIP string   `xml:"spbill_create_ip" json:"spbill_create_ip"`
}
