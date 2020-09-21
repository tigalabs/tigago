package tghttp

// message 消息提示
var message = make(map[string]map[string]string)

func init() {
	message["zh-CN"] = make(map[string]string)
	message["en-US"] = make(map[string]string)

	// PATH_NOT_FOUND
	message["zh-CN"]["PATH_NOT_FOUND"] = "路径找不到"
	message["en-US"]["PATH_NOT_FOUND"] = "PATH NOT FOUND"
}
