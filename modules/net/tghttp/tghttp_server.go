package tghttp

import (
	"fmt"
	"net/http"
)

// Server http server
type Server struct {
	router *router
	routes map[string]HandlerFunc // 路由映射 map
	lang   string                 // 语言
}

// NewServer 是 tghttp.Server 的构造函数
func NewServer() *Server {
	return &Server{
		router: newRouter(),
		lang:   "zh-CN", // 默认为中文
	}
}

// SetLang 设置语言
func (s *Server) SetLang(lang string) {
	s.lang = lang
	fmt.Printf("%#v", s)
}

// GET 定义 GET 请求方法
func (s *Server) GET(pattern string, handler HandlerFunc) {
	s.router.addRoute("GET", pattern, handler)
}

// POST 定义 POST 请求方法
func (s *Server) POST(pattern string, handler HandlerFunc) {
	s.router.addRoute("POST", pattern, handler)
}

// PUT 定义 PUT 请求方法
func (s *Server) PUT(pattern string, handler HandlerFunc) {
	s.router.addRoute("PUT", pattern, handler)
}

// DELETE 定义 DELETE 请求方法
func (s *Server) DELETE(pattern string, handler HandlerFunc) {
	s.router.addRoute("DELETE", pattern, handler)
}

// Run 运行自定义的http服务器的方法
func (s *Server) Run(addr string) (err error) {
	return http.ListenAndServe(addr, s)
}

// ServeHTTP 实现http.Handler接口
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	key := req.Method + "-" + req.URL.Path

	// 遍历匹配对应路由
	if handler, ok := s.router.routes[key]; ok {
		// 执行处理函数
		handler(newContext(w, req))
	} else {
		// 如果没有匹配到路由处理函数，则返回 404
		fmt.Fprintf(w, "PATH_NOT_FOUND: %s\n", req.URL)
	}
}
