package main

import "net/http"

/*
* 初衷: 使用多个Handler,处理各个不同的路由请求
* 1.不指定 Server struct里面的Handler 字段值, 默认就会走 DefaultServerMux
* 2.使用http.Handle 将某个的Handler 注册到 DefaultServerMux 上 (关系注意:)
*    a.http 包有一个 Handle函数
*    b. ServerMux struct (DefaultServeMux是一个ServeMux 指针变量) 也有一个Handle方法
*3.如果你调用 http.Handle 实际上就是调用 DefaultServeMux 上的Handle方法!!!

* 处理各个不同的路由请求 方法一:
* 使用 http.Handle 方法:
	type Handler interface {
		ServeHTTP(ResponseWriter,*Requset)
	}
  使用 http.HandleFunc 方法:
	Handle函数就是那些行为与Handle类似的函数,
	它的签名与 ServeHTTP方法的签名一致,接受参数
   Go有一种函数类型: HandleFunc. 可以将某个具有适当签名的函数f,适配成为一个Handler,而这个Handler具有方法f
*/

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About"))
}

func customHandleFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is customHandleFunc"))
}

func otherFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is otherFunc"))
}

func main() {

	mh := myHandler{}

	about := aboutHandler{}

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: nil, // 为nil 走 DefaultServeMux逻辑
	}

	// 使用 http.Handle 向DefaultServeMux 进行注册,
	//参数: 1.str 2:Handler指针类型
	http.Handle("/hello", &mh)

	// 注册第二个Handler
	http.Handle("/about", &about)

	// 直接用HandleFunc 注册 handler
	http.HandleFunc("/handle", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is HandleFunc"))
	})

	// 提取函数 再注册
	http.HandleFunc("/handleFunc", customHandleFunc)

	//  Handle + HandleFunc 通过go语言组合的形式 把其他的函数 转成Handle的形式
	//  http.HandlerFunc 他不是一个方法 是一个函数类型, 本身就是一个Handler,实现了Handler的一个接口,
	// 本质 http.HandlerFunc(xxx) 是类型转换
	http.Handle("/hf", http.HandlerFunc(otherFunc))

	// end
	server.ListenAndServe()

}

// 总结:
/*
*     http.Handle              http.HandleFunc
*     第二个参数是 Handler       第二个参数是 函数
*                              HandlerFunc,可以把Handler函数转化为Handler
							   内部调用还是用 http.Handle
*/
