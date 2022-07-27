package main

import (
	"math/rand"
	"net/http"
	"text/template"
	"time"
)

/*
Action 就是 Go 模板中嵌入的命令，位于两组花括号之间 {{ xxx }}
.就是一个 Action，而且是最重要的一个。它代表了传入模板的数据
·Action 主要可以分为五类:
	1.条件类 (arg 变量为true 就执行)
		{{ if arg }}
		 some content
		{{ end }}
	   变体:
	   	{{ if arg }}
		 some content
		{{ else }}
			other content
		{{ end }}
	2.迭代/遍历类
		{{ range array }}
		Dot is set to the element {{.}}
		{{ end }}
	.这类Action 用来遍历数组、slice、map 或 channel 等数据结构, “.”用来表示每次迭代循环中的元素
	回落机制: 当 range 后面为空 走 回落


	3.设置类
	{{ with arg }}
	 Dot is set to arg
	{{ end }}
	它允许在指定范围内 end结尾 ,让 “.”来表示其他指定的值
	4.包含类
	模版里包含模版, name 为模版名字(路径), value 为传入的动态值
	{{ template "name" arg }}

	5.定义类
	define action (暂时不说)
*/
func main() {
	server := http.Server{
		Addr: "localhost:8082",
	}

	http.HandleFunc("/process", process)

	http.HandleFunc("/range", handleRange)

	http.HandleFunc("/with", handleWith)

	http.HandleFunc("/empty", handleWithEmpty)

	http.HandleFunc("/tpl", handleTemplete)

	server.ListenAndServe()
}

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/tpl_if.html")
	// 随机数种子
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func handleRange(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/tpl_range.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func handleWith(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/tpl_with.html")
	t.Execute(w, "Hello")
}

func handleWithEmpty(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/tpl_with_empty.html")
	t.Execute(w, "Hello")
}

func handleTemplete(w http.ResponseWriter, r *http.Request) {
	tpl := "./wwwroot/"

	t, _ := template.ParseFiles(tpl+"template1.html", tpl+"template2.html")
	t.Execute(w, "hello world")

}
