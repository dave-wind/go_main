package main

import "net/http"

/*
*@创建 Web Server 方法二:
* http.Server 这是一个 struct
* Addr 字段表示网络地址
* Handler 字段: 如果为 nil 那么就是 DefaultServeMux
* ListenAndServe() 函数
 */
func main() {
	server := http.Server{
		Addr:    "localhost:8082",
		Handler: nil,
	}
	server.ListenAndServe()

	// 测试: 只生效一个 以下的不生效
	sever2 := http.Server{
		Addr:    "localhost:8083",
		Handler: nil,
	}
	sever2.ListenAndServe()
}

/*
* http.DefaultServeMux 作用:
*                           (会创建一个 goroutine 处理请求的过程就是用goroutine来完成)
*                                 Handler 1
*    http ===>  DefaultServeMux   Handler 2
*				 				  Handler 3
* 1. 多路复用器作用
* 2. 把进来的 http 请求 分配给 不同的 Handler, 相应的每个Handler 会在路由器上 进行注册
 */
