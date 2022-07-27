package main

import (
	"net/http"
	"text/template"
)

// Go 模版引擎工作原理
/*
1. client Request 请求来
2. Multiplexer 多路复用器 分配路由到对应的hanlder
3. 由handler来触发模版引擎, (template engine); Handler调用 Model
4. 模版引擎会去调用对应的模版( 通常是一组模版文件和动态数据 )
5. 模版引擎会生成HTML,并将其写入到ResponseWriter
6. ResponseWriter,再将它加入到HTTP响应中,返回给客户端
*/

// 关于模版
/*
内嵌的一些命令 (叫做 action)
text/template 是通用模版引擎, html/template 是 HTML模版引擎
action 位于 双层花括号之间: {{.}} ,可以命令模版引擎将其替换掉
*/

// 使用模版引擎
/*
* 1.解析模版源(可以是字符串或者模版文件),从而创建一个解析好的模版struct
* 2. 执行解析好的模版,并传入 ResponseWriter 和 数据
* 3.写入到 ResponseWriter 并返回
 */
func main() {
	server := http.Server{
		Addr: "localhost:8081",
	}

	http.Handle("/process", http.HandlerFunc(process))

	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	// 解析模版
	// 文件夹是相对路径, 比如 运行 go run . 就是当前执行的二进制文件的 相对路径
	t, _ := template.ParseFiles("./wwwroot/tpl.html")
	t.Execute(w, "Dave Wind")
}
