package core

func APIBaseURL() string {
	// TODO(samhuangszu): 后期做容灾功能
	return "https://api.mch.weixin.qq.com"
}

func APIRiskURL() string {
	return "https://fraud.mch.weixin.qq.com"
}
