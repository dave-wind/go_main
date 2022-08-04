package main

import (
	"net/http"
	"route/controller"
)

/*
1. 项目执行:
go mod init route
go run .


*/

// 路由结构

/*
                       home
					   handler

route ---> 前置------>  company
		 controller	   handler

					   about
					   handler

*/
func main() {
	// 导入写的模块
	controller.RegisterRoutes()
	// http.ListenAndServe("localhost:8081", nil)

	// https 走的是TLS层 https2.0版本
	http.ListenAndServeTLS("localhost:8081", "cert.pem", "key.pem", nil)
}

/*
开启https
先找到go的位置:
which go:
/usr/local/go/bin/go

然后在目标文件执行:                                     //设置host为localhost
go run /usr/local/go/src/crypto/tls/generate_cert.go -host localhost

生成文件和证书
*/

// 第三方路由器:
/*
gorilla/mux 灵活性高,功能强大,性能相对差
httprouter: 注重性能,功能简单
*/

/*
注意:
项目初始化 需要 go mod init xxx
自己写的模块 用package 包声明 可在内部使用, import 导入

*/
