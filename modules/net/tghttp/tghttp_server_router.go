package tghttp

// HandlerFunc 定义 tghttp 使用的请求处理程序
type HandlerFunc func(*Context)

type router struct {
	routes map[string]HandlerFunc // 路由映射 map
}

// newRouter 新建路由对象
func newRouter() *router {
	return &router{
		routes: make(map[string]HandlerFunc),
	}
}

// addRoute 新增路由
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.routes[key] = handler
}
