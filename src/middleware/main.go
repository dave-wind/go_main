package main

import (
	"encoding/json"
	"middleware/middleware"
	"net/http"
)

/*
---->  请求
                            ---->
				  中间件              Handler
				Middleware  <----

<----- 响应

请求时, 会优先走中间件, 响应也一样


2. 创建中间件
·func ListenAndServe(addr string, handler Handler) error

  handler 如果是 nil:DefaultServeMux

  不填nil 自己实现满足 ServeHTTP 接口的 Handler
· type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
  }




 3.中间件的用途

 Logging
 安全 (身份认证)
 请求超时
 响应压缩
*/

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}

		// 测试超时中间件时 取消注释:
		// time.Sleep(4 * time.Second)

		enc := json.NewEncoder(w)
		enc.Encode(c)
	})
	// 中间件 简单的嵌套关系: 先走timeout中间件,再走 Auth中间件
	http.ListenAndServe("localhost:8080", &middleware.TimeoutMiddleware{
		// new 一般返回一个类型的指针
		Next: new(middleware.AuthMiddleware),
	})
}
