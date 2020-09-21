package tghttp

var Message = make(map[string]map[string]string)

func init() {
	Message["zh-CN"] = make(map[string]string)
	Message["en-US"] = make(map[string]string)

	// PATH_NOT_FOUND
	Message["zh-CN"]["PATH_NOT_FOUND"] = "路径找不到"
	Message["en-US"]["PATH_NOT_FOUND"] = "PATH NOT FOUND"
}
