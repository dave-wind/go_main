package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

// 模版解析 :
/*
ParseFiles:
	解析模版文件, 并创建一个解析好的 struct,后续可以被执行
	ParseFiles函数是 Template struct 上的 一个ParseFiles 方法的简单调用
ParseFiles原理:
	调用ParseFiles后,会把文件解析为字符串, 再创建一个新的模版,模版名称是文件名, 最后调用Parse方法来解析字符串为模版
ParseFiles注意:
	参数可变,但只返回一个模版,
	当解析多个文件的时候,第一个文件作为返回的模版(名,内容),其余的作为map,供后续执行使用


* ParseGlob // 按模式匹配文件解析

* Parse // 解析字符串模版,其他方式最终都会调用Parse

* Lookup 方法 // 通过模版名来寻找模版,如果没找到就返回nil

* Must 函数 // 包裹一个函数,返回到一个模版的指针,和一个错误
           // 如果错误不为nil, 那么就 panic
		   // 作用: 主要用来减少解析模版错误

* 如何执行模版:
	Execute:
		参数是ResponseWriter,数据
		优点: 单模版,很适用
    ExecuteTemplate:
		参数是: ResponseWriter,模版名,数据
		模版集:	适用于多个模版

*/
func main() {

	server := http.Server{
		Addr: "localhost:8081",
	}

	http.Handle("/tmp1", http.HandlerFunc(tmp1))
	http.Handle("/tmp2", http.HandlerFunc(tmp2))
	http.Handle("/glob", http.HandlerFunc(HandleParseGlob))
	http.Handle("/moretpl", http.HandlerFunc(HandleExecuteTemplate))

	server.ListenAndServe()

}

func tmp1(w http.ResponseWriter, r *http.Request) {
	// 效果 如上一样, 先创建空的模版, 在解析模版填充
	t, _ := template.ParseFiles("./wwwroot/tpl.html")

	t.Execute(w, "Dave Wind1")
}

func tmp2(w http.ResponseWriter, r *http.Request) {
	// 效果 如上一样, 先创建空的模版, 在解析模版填充
	t2 := template.New("./wwwroot/tpl.html")
	t2, _ = t2.ParseFiles("./wwwroot/tpl.html")

	t2.Execute(w, "Dave Wind2")
}

type Poster struct {
	name string
}

func HandleParseGlob(w http.ResponseWriter, r *http.Request) {
	// 返回的是找到的第一文件名
	t, _ := template.ParseGlob("./wwwroot/*.html")

	p := &Poster{
		name: "dave",
	}
	fmt.Fprintln(w, t.Name())
	fmt.Fprintln(w, p.name)
}

func HandleExecuteTemplate(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("./wwwroot/tpl.html")
	// t.Execute(w, "t")

	// 填多个参数 为文件相对地址
	ts, err := template.ParseFiles("./wwwroot/tpl.html", "./wwwroot/tpl2.html")
	if err != nil {
		log.Println("err", err)
	}
	// 第二个参数就是 文件名而已
	ts.ExecuteTemplate(w, "tpl2.html", "指定的是第二个模版")

}
