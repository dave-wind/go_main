package main

import (
	"fmt"
	"net/http"
)

/* MultipartForm
* 想要使用 MultipartForm 这个字段的话，首先需要调用 ParseMultipartForm 这个方法
。该方法会在必要时调用 ParseForm 方法·参数是需要读取数据的长度
	·MultipartForm 只包含表单的 key-value 对
	·返回类型是一个 struct 而不是 map。这个 struct 里有两个 map;
	· key 是 string,value 是 []string
	。空的(key 是 string，value 是文件)
*/

func main() {
	server := http.Server{
		Addr: "localhost:8081",
	}

	// MultipartForm
	http.HandleFunc("/multipart", func(w http.ResponseWriter, r *http.Request) {
		// 解析 multipart/form-data form 需要用该方法 并传入字节
		r.ParseMultipartForm(1024)
		// 解析完 后访问拿 MultipartForm
		fmt.Fprintln(w, r.MultipartForm) // &{map[first_name:[111] last_name:[222]] map[]}
	})

	// FormValue 和 PostFormValue
	http.Handle("/form", http.HandlerFunc(HandleForm))

	server.ListenAndServe()
}

/*
·FormValue 方法会返回 Form 字段中指定 key 对应的第一个 value
·无需调用 ParseForm 或 ParseMultipartForm
、PostFormValue 方法也一样，但只能读取 PostForm
·FormValue 和 PostFormValue 都会调用 ParseMultipartForm 方法。
但如果表单的 enctype 设为 multipart/form-data，那么即使你调用 ParseMultipartForm 方法，也无法通过 FormValue 获得想要的值。
*/
func HandleForm(w http.ResponseWriter, r *http.Request) {
	// 如果 enctype 为 application/x-www-form-urlencoded, 可以这么拿第一个值
	// 	// 如果 enctype 为 multipart/form-data, 拿到的是 查询字符的第一个
	// r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.FormValue("first_name"))

	// 只读表单的值
	fmt.Fprintln(w, r.PostFormValue("first_name"))

}

// 注意: enctype: multipart/form-data 尽量避免使用FormValue
