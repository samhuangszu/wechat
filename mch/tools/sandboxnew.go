package tools

import (
	"errors"

	"github.com/samhuangszu/wechat/mch/core"
)

// GetSignKey 获取沙箱paykey
func GetSignKey(clt *core.Client, req map[string]string) (string, error) {
	resp, err := clt.PostXML("https://api.mch.weixin.qq.com/sandboxnew/pay/getsignkey", req)
	if err != nil {
		return "", err
	}
	//获取signkey，返回
	key, ok := resp["sandbox_signkey"]
	if !ok {
		return "", errors.New("sandbox_signkey not exist")
	}
	if len(key) <= 0 {
		return "", errors.New("sandbox_signkey not exist")
	}
	return key, nil
}

// GetSignKeyReq 获取沙箱paykey
type GetSignKeyReq struct {
	XMLName  struct{} `xml:"xml" json:"-"`
	MchID    string   `xml:"mch_id" json:"mch_id"`
	NonceStr string   `xml:"nonce_str" json:"nonce_str"`
	Sign     string   `xml:"sign" json:"sign"`
}
