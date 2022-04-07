package main

import "net/http"

/*
* handler 是一个接口 (interface)
* handler 定义了一个方法 ServeHTTP(),参数:
* 1.HTTPResponseWriter
* 2.指向Request 这个struct 的指针
 */
type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}
func main() {
	mh := myHandler{}

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: &mh,
	}
	server.ListenAndServe()

}
