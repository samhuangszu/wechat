package mmpaymkttransfers

import (
	"github.com/samhuangszu/wechat/mch/core"
)

// 发放裂变红包.
//  NOTE: 请求需要双向证书
func SendGroupRedPack(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+core.GetPath(clt, "/mmpaymkttransfers/sendgroupredpack"), req)
}
