package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
* 1.multipart/form-data 最常见的应用场景就是上传文件
*   首先调用 ParseMultipartForm 方法
*   从 File字段获得 FileHeader,调用其Open方法来获得文件
*   可以使用 ioutil.ReadAll 函数把文件内容读取到 byte切片里
* 2.FormFile方法
*   无需调用 ParseMultipartForm 方法返回指定 key 对应的第一个 value
*   同时返回 File 和 FileHeader，以及错误信息
*   如果只上传一个文件, 这种方法就很快捷
 */
func main() {

	server := http.Server{
		Addr: "localhost:8081",
	}

	http.Handle("/process", http.HandlerFunc(process))

	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {

	// 方法一:
	// r.ParseMultipartForm(1024)

	// // 拿到fileHeader的指针
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()

	// 方法二: 用 request.FormFile 方法,就无须调用 ParseMultipartForm方法

	// uploaded 字段用的是form表单的字段, 返回的是对应的第一个文件, _:为 fileHeader
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		// 读取为 byte slice切片
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

// 注意:
/*
* 不是所有的 POST 请求都来自 Form
· 客户端框架(例如 Angular 等) 会以不同的方式对 POST 请求编码:
·jQuery 通常使用 application/x-w www-form-urlencoded
*.Angular 是 application/json
 ParseForm 方法无法处理 application/json
*/

// 难点: MultipartReader 流的方式处理

/*
* 1.MultipartReader 的签名: ·func (r *Request) MultipartReader() (*multipart.Reader, error)

* 2.如果是 multipart/form-data 或 multipart 混合的 POST 请求:
·MultipartReader 返回一个 MIME multipart reader。否则返回 nil 和一个错误
·可以使用该函数代替 ParseMultipartForm 来把请求的 body 作为 stream 进行处理
 特点: 不是把表单作为一个对象来处理的，不是一次性获得整个 map·逐个检查来自表单的值，然后每次处理一个
*/
