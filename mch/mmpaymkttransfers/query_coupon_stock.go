package mmpaymkttransfers

import (
	"github.com/samhuangszu/wechat/mch/core"
)

// 查询代金券批次信息.
func QueryCouponStock(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+core.GetPath(clt, "/mmpaymkttransfers/query_coupon_stock"), req)
}
