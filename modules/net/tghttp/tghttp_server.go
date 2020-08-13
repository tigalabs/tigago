package tghttp

import (
	"fmt"
	"net/http"
)

// HandlerFunc定义tghttp使用的请求处理程序

type HandlerFunc func(*Context)

type Server struct {
}

// TODO:定义新增路由 路由分组方法

// Run 运行自定义的http服务器的方法
func (server *Server) Run(addr string) (err error) {
	return http.ListenAndServe(addr, server)
}

// ServeHTTP 实现http.Handler接口
func (server *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	fmt.Print(c)
	// server.router.handle(c)
}
