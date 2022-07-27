package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8082",
	}

	http.Handle("/responsewriter", http.HandlerFunc(HandleWrite))
	http.HandleFunc("/writeheader", HandleWriteHeader)
	http.Handle("/redirect", http.HandlerFunc(WriteSetHeader))
	http.Handle("/json", http.HandlerFunc(HandleJson))
	server.ListenAndServe()
}

// 关于 Response:
/*
* 为什么 Handler的 ServeHTTP,只有一个
 */
// *http.Request 是按引用进行传递, 改指针 就会改动他的值
// http.ResponseWriter 本身是一个接口, response 指针实现了 ResponseWriter接口, 所以 ResponseWriter就代表着 response指针

// 写入ResponseWriter
/*
* Write方法接收一个byte切片作为参数,然后把它写入到HTTP响应到Body 里面
* 如果在Write方法被调用时, header里面 没有设定 content typ, 那么数据的前512个字节就会被检测 content type
 */
func HandleWrite(w http.ResponseWriter, r *http.Request) {
	// str := `<html>
	// <head><title>Go Web</title></head>
	// <body><h1>Hello World</h1></body>
	// </html>`
	str := `123456789abcdefghijklmnopqrstuvwxyz`
	w.Write([]byte(str))
}

// 运行:
// 1  go run xxx.go
// 2  curl -i localhost:8082/responsewriter

// 结果:
/*
* HTTP/1.1 200 OK
* Date: Sat, 09 Jul 2022 02:40:28 GMT
* Content-Length: 86
* Content-Type: text/html; charset=utf-8
 */

// Write Header 方法
// WriteHeader 方法接收一个整数(HTTP状态码)作为参数,并把它作为HTTP响应的状态码返回
// 如果该方法 没有显式调用,那么在第一次调用Write方法前,会隐式的调用 WriteHeader(http.StatusOK)
// 所以 WriteHeader 主要是用来发送错误类的 HTTP状态码
// 调用完 WriteHeader 方法之后,仍然可以写入到ResponseWriter,但无法再修改 header了
func HandleWriteHeader(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

// 运行: // 2  curl -i localhost:8082/writeheader

func WriteSetHeader(w http.ResponseWriter, r *http.Request) {
	// Header方法 返回一个headers map ,可以进行修改, (需要在WriteHeader之前进行修改)
	// 修改后的headers 将会体现在返回给客户端的 HTTP 响应里
	// w.Header() 返回一个 map http header
	w.Header().Set("Location", "http://www.baidu.com")

	// 跳转
	w.WriteHeader(302)
}

type Post struct {
	User    string
	Threads []string
}

func HandleJson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &Post{
		User:    "Dave Wind",
		Threads: []string{"first", "second", "thrid"},
	}
	// json.Marshal: 把struct字面值 转为 json
	json, _ := json.Marshal(post)
	w.Write(json)
}
