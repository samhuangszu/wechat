package mmpaymkttransfers

import (
	"github.com/samhuangszu/wechat/mch/core"
)

//GetTransferInfo 查询企业付款.
//  NOTE: 请求需要双向证书
func GetTransferInfo(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/gettransferinfo", req)
}

// TransferInfoReq 查询付款到银行卡
type TransferInfoReq struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	MchID          string   `xml:"mch_id" json:"mch_id"`
	AppID          string   `xml:"appid" json:"appid"`
	PartnerTradeNo string   `xml:"partner_trade_no" json:"partner_trade_no"`
	NonceStr       string   `xml:"nonce_str" json:"nonce_str"`
	Sign           string   `xml:"sign" json:"sign"`
}
