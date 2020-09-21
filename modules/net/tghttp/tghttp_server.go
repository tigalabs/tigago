package tghttp

import (
	"fmt"
	"net/http"
)

// HandlerFunc定义tghttp使用的请求处理程序
type HandlerFunc func(*Context)

type Server struct {
	// 路由映射 map
	routes map[string]HandlerFunc
	// 语言
	lang string
}

// TODO:定义新增路由 路由分组方法

// NewServer 是 tghttp.Server 的构造函数
func NewServer() *Server {
	return &Server{
		routes: make(map[string]HandlerFunc), // 初始化路由
		lang:   "zh-CN",                      // 默认为中文
	}

}

// SetLang 设置语言
func (server *Server) SetLang(lang string) {
	server.lang = lang
	fmt.Printf("%#v", server)
}

// addRoute 新增路由
func (server *Server) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	server.routes[key] = handler
}

// GET 定义 GET 请求方法
func (server *Server) GET(pattern string, handler HandlerFunc) {
	server.addRoute("GET", pattern, handler)
}

// POST 定义 POST 请求方法
func (server *Server) POST(pattern string, handler HandlerFunc) {
	server.addRoute("POST", pattern, handler)
}

// Run 运行自定义的http服务器的方法
func (server *Server) Run(addr string) (err error) {
	return http.ListenAndServe(addr, server)
}

// ServeHTTP 实现http.Handler接口
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	key := req.Method + "-" + req.URL.Path

	// 遍历匹配对应路由
	if handler, ok := server.routes[key]; ok {
		// 执行处理函数
		handler(newContext(w, req))
	} else {
		// 如果没有匹配到路由处理函数，则返回 404
		fmt.Fprintf(w, Message[server.lang]["PATH_NOT_FOUND"]+": %s\n", req.URL)
	}
}
