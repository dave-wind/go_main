package main

import "net/http"

/*
*@description: 内置 Handlers:

* NotFoundHandler
	func NotFoundHandler() Handler

* RedirectHandler
	func RedirectHandler(url string,code int) Handler
	返回一个 Handler,它把每个请求使用给定的状态码跳转到指定的URL
	url: 要跳转的URL
	code: 跳转的状态码(3xx),常见的: StatusMovedPermanently,StatusFound或 StatusSeeOther等

* StripPrefix
	func StripPrefix(prefix string, h handler) Handler
	返回一个handler,它从请求URL中去掉指定的前缀,然后再调用另一个Handler
	  1.如果请求的URL与提供的前缀不符,那么404
    略像中间件
		1.prefix,URL将要被移除的字符串前缀
		2.h,是一个handler,在移除字符串前缀之后,这个hanlder将会接收到请求
	修饰了另外一个Handler

* TimoutHander
  function TimeoutHandler(h Handler,dt time.Duration,msg string) Handler
  返回一个handler,它用来在指定时间内运行传入的h
  也相当于是一个修饰器
	h,将要被修饰的handler
	dt,第一个handler允许的处理时间
	msg,如果超时,那么就把msg返回给请求,表示响应时间过长

* FileServer
* func FileServer(root FileSystem) Handler
* 返回一个 handler,使用基于root文件系统来响应请求
* type FileSystem interface {
	Open(name string)(File, error)
}
* 使用时需要用到操作系统的文件系统,所以还需要委托给:
* type Dir string
* func (d Dir) Open(name string)(File, error)
*/

func demo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Demo"))
}

func detailFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("1"))
}

func main() {

	http.Handle("/demo", http.HandlerFunc(demo))

	// NotFoundHandler
	http.Handle("/404", http.NotFoundHandler())

	// RedirectHandler
	http.Handle("/demo2", http.RedirectHandler("/404", 404))

	// ??
	http.Handle("/detail", http.StripPrefix("detail", http.HandlerFunc(detailFunc)))

	// FileServer ?
	http.Handle("/", http.FileServer(http.Dir("wwwroot")))

	// 端口为 localhost
	// nil : 使用 DefaultServeMux 路由器 进行分配
	http.ListenAndServe(":8083", nil)
	// http.ListenAndServe(":8083", http.FileServer(http.Dir("wwwroot")))

}

// 疑问:
/*
* 1.go build . && web
* 2.StripPrefix
* 3.TimoutHander
* 4.FileServer 有路由就不展示,根目录就可以
 */
