package mmpaysptrans

import "github.com/samhuangszu/wechat/mch/core"

// QueryBankOrder 查询企业付款.
func QueryBankOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaysptrans/query_bank", req)
}

// QueryBankRequest 查询付款到银行卡
type QueryBankRequest struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	MchID          string   `xml:"mch_id"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
}

// GetPublicKey 获取Rsa公钥
func GetPublicKey(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/risk/getpublickey", req)
}

// PublicKeyReq 查询付款到银行卡
type PublicKeyReq struct {
	XMLName  struct{} `xml:"xml" json:"-"`
	MchID    string   `xml:"mch_id"`
	NonceStr string   `xml:"nonce_str"`
	Sign     string   `xml:"sign"`
	SignType string   `xml:"sign_type"`
}

// Paybank 付款到银行卡
func Paybank(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaysptrans/pay_bank", req)
}

// PaybankReq 查询付款到银行卡
type PaybankReq struct {
	XMLName        struct{} `xml:"xml" json:"-"`
	MchID          string   `xml:"mch_id"`
	PartnerTradeNo string   `xml:"partner_trade_no"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	EncBankNo      string   `xml:"enc_bank_no"`
	EncTrueName    string   `xml:"enc_true_name"`
	BankCode       string   `xml:"bank_code"`
	Amount         int      `xml:"amount"`
	Desc           string   `xml:"desc"`
}
