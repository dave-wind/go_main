package main

import (
	"net/http"
	"text/template"
)

/*
管道是按顺序连接到一起的参数、函数和方法。·和 Unix 的管道类似
· 例如:{{ p1 | p2 | p3 }}
·p1、p2、p3 要么是参数，要么是函数
管道允许我们把参数的输出发给下一个参数，下一个参数由管道(|)分隔开。
*/
func main() {
	server := http.Server{
		Addr: "localhost:8082",
	}

	http.HandleFunc("/process", handlePipeline)

	server.ListenAndServe()
}

func handlePipeline(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./wwwroot/pipeline.html")
	t.Execute(w, 3.1415926)
}
