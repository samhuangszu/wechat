package core

import "strings"

func APIBaseURL() string {
	// TODO(samhuangszu): 后期做容灾功能
	return "https://api.mch.weixin.qq.com"
}

func APIRiskURL() string {
	return "https://fraud.mch.weixin.qq.com"
}

// if !(strings.Index(url, "risk/getpublickey") > -1 || strings.Index(url, "mmpaysptrans/pay_bank") > -1 || strings.Index(url, "mmpaysptrans/query_bank") > -1) {

var notAppIDURL = []string{
	"risk/getpublickey",
	"mmpaysptrans/pay_bank",
	"mmpaysptrans/query_bank",
}

// NotCheckRespAppID 不检验返回AppID
func NotCheckRespAppID(url string) bool {
	for _, u := range notAppIDURL {
		if strings.Index(url, u) >= 0 {
			return true
		}
	}
	return false
}

//不需要校验返回值的URL
var notSignURL = []string{
	"mmpaymkttransfers/promotion/transfers",
	"mmpaymkttransfers/gettransferinfo",
	"mmpaymkttransfers/sendredpack",
	"mmpaymkttransfers/sendgroupredpack",
	"mmpaymkttransfers/gethbinfo",
	"risk/getpublickey",
	"mmpaysptrans/query_bank",
}

// NotCheckRespSign 不检验返回签名
func NotCheckRespSign(url string) bool {
	for _, u := range notSignURL {
		if strings.Index(url, u) >= 0 {
			return true
		}
	}
	return false
}

//不需要把ID传到req中
var notNeedURL = []string{
	"mmpaymkttransfers/promotion/transfers",
	"mmpaymkttransfers/sendredpack",
	"mmpaymkttransfers/sendgroupredpack",
	"mmpaysptrans/pay_bank",
	"mmpaysptrans/query_bank",
}

// NotNeedIds 不需要把ID传到req中
func NotNeedIds(url string) bool {
	for _, u := range notNeedURL {
		if strings.Index(url, u) >= 0 {
			return true
		}
	}
	return false
}

// SandBoxURL 沙箱路径
var SandBoxURL = map[string]string{

	"/pay/micropay":                          "/sandboxnew/pay/micropay",
	"/pay/unifiedorder":                      "/sandboxnew/pay/unifiedorder",
	"/pay/closeorder":                        "/sandboxnew/pay/closeorder",
	"/pay/downloadbill":                      "/sandboxnew/pay/downloadbill",
	"/pay/orderquery":                        "/sandboxnew/pay/orderquery",
	"/secapi/pay/reverse":                    "/sandboxnew/secapi/pay/reverse",
	"/secapi/pay/refund":                     "/sandboxnew/pay/refund",
	"/pay/refundquery":                       "/sandboxnew/pay/refundquery",
	"/mmpaysptrans/pay_bank":                 "/sandboxnew/mmpaysptrans/pay_bank",
	"/mmpaysptrans/query_bank":               "/sandboxnew/mmpaysptrans/query_bank",
	"/risk/getpublickey":                     "/sandboxnew/risk/getpublickey",
	"/promotion/query_coupon":                "/sandboxnew/promotion/query_coupon",
	"/mmpaymkttransfers/promotion/transfers": "/sandboxnew/mmpaymkttransfers/promotion/transfers",
	"/mmpaymkttransfers/gethbinfo":           "/sandboxnew/mmpaymkttransfers/gethbinfo",
	"/mmpaymkttransfers/gettransferinfo":     "/sandboxnew/mmpaymkttransfers/gettransferinfo",
	"/mmpaymkttransfers/query_coupon_stock":  "/sandboxnew/mmpaymkttransfers/query_coupon_stock",
	"/mmpaymkttransfers/send_coupon":         "/sandboxnew/mmpaymkttransfers/send_coupon",
	"/mmpaymkttransfers/sendgroupredpack":    "/sandboxnew/mmpaymkttransfers/sendgroupredpack",
	"/mmpaymkttransfers/sendredpack":         "/sandboxnew/mmpaymkttransfers/sendredpack",
}

// ProductURL 正在环境路径
var ProductURL = map[string]string{
	"/pay/micropay":                          "/pay/micropay",
	"/pay/unifiedorder":                      "/pay/unifiedorder",
	"/pay/closeorder":                        "/pay/closeorder",
	"/pay/orderquery":                        "/pay/orderquery",
	"/pay/downloadbill":                      "/pay/downloadbill",
	"/secapi/pay/reverse":                    "/secapi/pay/reverse",
	"/secapi/pay/refund":                     "/secapi/pay/refund",
	"/pay/refundquery":                       "/pay/refundquery",
	"/mmpaysptrans/query_bank":               "/mmpaysptrans/query_bank",
	"/risk/getpublickey":                     "/risk/getpublickey",
	"/promotion/query_coupon":                "/promotion/query_coupon",
	"/mmpaysptrans/pay_bank":                 "/mmpaysptrans/pay_bank",
	"/mmpaymkttransfers/promotion/transfers": "/mmpaymkttransfers/promotion/transfers",
	"/mmpaymkttransfers/gethbinfo":           "/mmpaymkttransfers/gethbinfo",
	"/mmpaymkttransfers/gettransferinfo":     "/mmpaymkttransfers/gettransferinfo",
	"/mmpaymkttransfers/query_coupon_stock":  "/mmpaymkttransfers/query_coupon_stock",
	"/mmpaymkttransfers/send_coupon":         "/mmpaymkttransfers/send_coupon",
	"/mmpaymkttransfers/sendgroupredpack":    "/mmpaymkttransfers/sendgroupredpack",
	"/mmpaymkttransfers/sendredpack":         "/mmpaymkttransfers/sendredpack",
}

// GetPath 获取沙箱或生产环境的路径
func GetPath(clt *Client, path string) string {
	if clt.SandBox() {
		url, ok := SandBoxURL[path]
		if ok {
			return url
		}
		return path
	}
	url, ok := ProductURL[path]
	if ok {
		return url
	}
	return path
}
