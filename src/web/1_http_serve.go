package main

import "net/http"

/*
*@创建 Web Server 方法一:
* http.ListenAndServe()
* 第一个参数 是网络地址, 如果为“” 那么就代表所有的网络接口的80端口
* 第二个参数 是handler, 如果为nil, 那么就是 DefaultServeMux
* DefaultServeMux: 是一个multiplexer (可以看作是路由器)
 */
func main() {
	// 没有对具体的网址 设定路由 (提供特定的具体逻辑)
	// 运行 也是 404
	http.ListenAndServe("localhost:8081", nil)
}
