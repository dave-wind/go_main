package main

import (
	"net/http"
	"text/template"
	"time"
)

/*
内置函数

·define、template、block
·html、js、urlquery。对字符串进行转义，防止安全问题
。如果是 Web 模板，那么不会需要经常使用这些函数。
·index
·print/printf/println
·len
·with
*/

/*
如何自定义函数
了解1:
template.Funcs(funcMap FuncMap)*Template   //返回 一个Template
了解2:
type FuncMap map[string]interface{}  // 表示 map集合 key为string, value是空接口,可以是任意类型
	value 是函数
		可以任意数量的参数
		返回单个值的函数或返回一个值+一个错误函数
步骤1:
 创建 一个 FuncMap, key为函数别名, value就是函数
步骤2:
 把FuncMap 附加到模版上


*/

/*
如何使用自定义函数
 常见方法:  templete.New("").Funcs(funcMap).Parse(...)
 调用顺序非常重要
*/

func main() {
	server := http.Server{
		Addr: "localhost:8083",
	}

	http.HandleFunc("/func", handleFunc)
	server.ListenAndServe()
}

func handleFunc(w http.ResponseWriter, r *http.Request) {
	tpl := "./wwwroot/tpl_func.html"
	// 创建map
	funcMap := template.FuncMap{"fdate": formatDate}

	// 附加模版
	t := template.New("tpl_func.html").Funcs(funcMap)

	// 	解析模版
	t.ParseFiles(tpl)
	t.Execute(w, time.Now())
}

func formatDate(t time.Time) string {
	// 格式化
	layout := "2006-01-02"
	return t.Format(layout)
}
