package promotion

import (
	"github.com/samhuangszu/wechat/mch/core"
)

// 查询代金券信息.
func QueryCoupon(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+core.GetPath(clt, "/promotion/query_coupon"), req)
}
