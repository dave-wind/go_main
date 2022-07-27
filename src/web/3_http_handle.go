package main

import "net/http"

/*
* 思考逻辑:
* 1.http.Server 可以创建 http请求
* 2.参数是一个Handler 是一个函数方法, 传的自定义handler 需要符合 handler的接口要求!
* 3.分析 handler
* handler 是一个接口 (interface)
* handler 定义了一个方法 ServeHTTP(),参数:
* 1.HTTPResponseWriter
* 2.指向Request 这个struct 的指针
 */

// 声明一个结构体类型, 再其上定义了一个 ServeHTTP方法, 那么在Go语言里,符合了一个接口的要求 那就是符合了接口的类型
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func main() {
	mh := myHandler{}

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: &mh, // 传指针
	}
	server.ListenAndServe()

}
