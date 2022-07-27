package main

import (
	"fmt"
	"net/http"
)

/* 1.Form内容:
	Form字段
	PostmForm字段
	MultipartForm字段
	FormValue & PostFormValue 方法
	上传文件(Files)
	POST JSON
 HTML表单里的数据会以 name-value对的形式,通过POST请求发送出去
* 2.表单post请求数据格式是什么样的?
  发送的数据格式是根据表单的Content Type来指定的,也就是enctype属性:
  默认: application/x-www-form-urlencoded
  浏览器至少支持: application/x-www-form-urlencoded, multipart/form-data
  如果是html5 还需要支持 text/plain
*
* 3.如何选择?
*  简单文本: 表单URL编码 提交
*  大量数据,例如上传文件: multipart-MIME (每一条name-value都会转换成一个MIME消息,以及自己的content-type和Content Disposition)
*    甚至可以把二进制数据通过 Base64编码,来当文本进行发送
*
* 4.Form 字段:
	Request 上的函数允许我们从 URL 或/和 Body 中提取数据，通过这些字段:
		·Form
		·PostForm
		·MultipartForm
	Form 里面的数据是 key-value 对。通常的做法是:
		·先调用 ParseForm 或 ParseMultipartForm 来解析 Request
		·然后相应的访问 Form、PostForm 或 MultipartForm 字段
*/

func main() {
	server := http.Server{
		Addr: "localhost:8081",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		// 先解析 request
		r.ParseForm()
		// 再拿 r.Form Form的底层数据类型 是一个 map[string][]string, 所有当有多个key 都会被收集起来
		fmt.Fprintln(w, r.Form)
		// 只拿 form字段,不会拿url上带的查询字符
		// fmt.Fprintln(w, r.PostForm)
		fmt.Fprintln(w, r.PostForm["first_name"])
	})

	server.ListenAndServe()
}

// !!!!!! ParseForm,PostForm 适用于 application/x-www-form-urlencoded
// 如果 form表单 action带有查询参数,只想拿form表单的 需要访问 r.postForm字段
/*
*  <form action="http://localhost:8081/process?first_name=Nick"
 */
