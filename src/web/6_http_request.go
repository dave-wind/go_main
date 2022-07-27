package main

import (
	"fmt"
	"log"
	"net/http"
)

/* Http 消息:
*  1. HTTP Request 和 HTTP Response (请求和响应)
*  2.他们具有相同的结构:
*	 请求(响应)行 (URL字段代表了它的部分内容)
*    0个或多个 Header
*    空行
*    可选的消息体(Body)
* Request(是一个struct),代表了客户端发送的HTTP请求信息
* 重要的字段:
* URL,Header,Body,Form,PostForm,MultipartForm
*
 */

/* 请求的URL
 * Request的URL字段,代表的是请求行
 * URL字段是指向的url.URL 类型的一个指针, url 是一个包, url.URL是一个struct
 * type URL struct {
	 Scheme string
	 Opaque string
	 User   *Userinfo
	 Host string
	 Path string
	 RawQuery string
	 Fragment string
 }
*/

// URL通用格式: scheme://[userinfo@]host/path[?query][#fragment]
// 不以斜杠开头的URL被解释为: shceme:opaque[?query][#fragment]  eg:baidu.com

// (RawQuery)如何拿到 URL查询字符串:
/** RawQuery
eg: https://www.baidu.com?id=1&name=dave
RawQuery的值就是 id=1&name=dave, 还可以通过form 拿到 key-value
*/

// (Fragment) URL Fragment
/*
如果浏览器里发出请求,那么你无法提取 Fragment字段的值
因为:浏览器发送请求 会把fragment部分去掉
(http://www.baidu.com/demo.htm#print)
*/

// Request Header
/*
* 请求和响应 (Request 和 Response)的headers 是通过Header类型来描述的
* 它是一个map, 用来表述 HTTP Header 里的 key-value对
* Header map 的key 是string类型, value 是 string[] 切片
* 设置key的时候会创建一个 空的 []string作为value, value里面第一个元素就是新的header 值
* 为指定的key添加一个新的Header值,执行append操作即可
 */

// Request Header 例子:

/* r: request
 *  r.Header (获得所有的Header, 一个map)
 *	r.Header[""] 返回 []string  多个值
 *  r.Header.Get("Accept-Encoding") 返回string 用逗号隔开
 */
func main() {
	server := http.Server{
		Addr: "localhost:8083",
	}
	// header demo
	http.HandleFunc("/header", func(w http.ResponseWriter, r *http.Request) {

		// fmt.Fprintln 可以把内容写入到 http.ResponseWriter 中
		fmt.Fprintln(w, r.Header)
		fmt.Fprintln(w, r.Header["Accept-Encoding"])
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding"))
		fmt.Fprintln(w, "html")
	})

	// body
	http.Handle("/body", http.HandlerFunc(HandleBody))

	// query
	http.Handle("/query", http.HandlerFunc(HandleQuery))

	server.ListenAndServe()
}

// Request 和 Body
/*
 * 请求和响应的bodies都是使用Body字段来表示的
 * Body是一个 io.ReadCloser接口 (Body的类型是 io.ReadCloser)
 *  一个Reader接口
 *  一个Closer接口
 * Reader接口定义了一个	Read方法:
 * 参数: []byte
 * 返回: byte的数量,可选的错误
 * Closer接口定义了一个Close方法
 * 没有参数,返回可选的错误
 */

func HandleBody(w http.ResponseWriter, r *http.Request) {
	// 首先获取内容长度
	length := r.ContentLength
	// 创建一个空的切片, 容量是length
	body := make([]byte, length)
	// 通过 r.Body Read方法 读取Body内容 到 body变量中
	result, er := r.Body.Read(body)
	// 最后写入到 ResponseWriter 返回
	if er != nil {
		fmt.Fprintln(w, string(body))
		fmt.Fprintln(w, er)
		fmt.Fprintln(w, result)
	}
}

// 查询参数(query Parameters)
/*
* 1.URL Query
*   eg: https://www.baidu.com/post?id=123&thread_id=456
*   r.URL.RawQuery 会提供实际查询的原始字符串
*   上例的 RawQuery的值就是 id =123&thread_id=456
*
* r.URL.Query(),会提供查询字符串对应的map[string][]string (key是字符串,值是 字符串切片)
 */

//  eg: 浏览器访问 http://localhost:8083/query?id=1&name=dave&id=2&name=wind
func HandleQuery(w http.ResponseWriter, r *http.Request) {
	url := r.URL

	query := url.Query() // 返回一个map

	id := query["id"]
	log.Println(id) // 返回的是一个切片 因为可能有多个

	name := query.Get("name")
	log.Println(name) // 返回字符串 第一个

	fmt.Fprintln(w, id)
	fmt.Fprintln(w, name)
}
