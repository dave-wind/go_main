package main

import (
	"log"
	"net/http"
	"text/template"
)

/*
* 如何制作layout模版
* ·Include(包含)action 的形式:{{ template "name".}}
* ·正确的做法是在模板文件里面使用 define action 再定义一个模板 见 layout
*  可以通过 define 在layout定义相同的模版名 动态替换
 */
/*

block模版
{ block arg }} Dot is set to arg{{ end }}
·block action 可以定义模板，
并同时就使用它·template:模板必须可用·block:模板可以不存在

and
    返回第一个为空的参数或最后一个参数。可以有任意多个参数。
    and x y等价于if x then y else x

not
    布尔取反。只能一个参数。

or
    返回第一个不为空的参数或最后一个参数。可以有任意多个参数。
    "or x y"等价于"if x then x else y"。

print
printf
println
    分别等价于fmt包中的Sprint、Sprintf、Sprintln

len
    返回参数的length。

index
    对可索引对象进行索引取值。第一个参数是索引对象，后面的参数是索引位。
    "index x 1 2 3"代表的是x[1][2][3]。
    可索引对象包括map、slice、array。

call
    显式调用函数。第一个参数必须是函数类型，且不是template中的函数，而是外部函数。
    例如一个struct中的某个字段是func类型的。
    "call .X.Y 1 2"表示调用dot.X.Y(1, 2)，Y必须是func类型，函数参数是1和2。
    函数必须只能有一个或2个返回值，如果有第二个返回值，则必须为error类型。
*/

/*
* 补充逻辑运算符
* 	 eq/ne : 相等/不相等
*    lt/gt:  小于/大于
*    le/ge: 小于等于/大于等于
*    and : 逻辑与
*    or: 逻辑或
*    not: 逻辑非
 */

func main() {

	http.Handle("/home", http.HandlerFunc(HanleHome))

	http.HandleFunc("/about", HanleAbout)

	http.HandleFunc("/struct", HandleStruct)

	http.ListenAndServe("localhost:8081", nil)
}

func HanleHome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/layout.html", "./wwwroot/home.html")
	e := t.ExecuteTemplate(w, "layout", "home")
	log.Println(e)
}

func HanleAbout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/layout.html", "./wwwroot/about.html")
	t.ExecuteTemplate(w, "layout", "Hello About")
}

type Person struct {
	Name string
	Age  int
}

func HandleStruct(w http.ResponseWriter, r *http.Request) {
	p := Person{"Dave", 28}
	// 解析模版
	t, _ := template.ParseFiles("./wwwroot/layout.html", "./wwwroot/struct.html")
	// 设置作用域值
	tpl, _ := t.Parse("Name:{{.Name}}, Age:{{.Age}}")
	tpl.ExecuteTemplate(w, "layout", p)
}
