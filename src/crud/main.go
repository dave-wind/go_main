package main

import (
	"context"
	"crud/common"
	"crud/controller"
	"crud/middleware"
	"database/sql"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	IP       = "127.0.0.1"
	PORT     = "3306"
	USERNAME = "root"
	PASSWORD = "1234qwer"
	dbName   = "dave"
)

// 每个包中可以有多个init函数，而其他包的init是在main包的init调用之前被执行，main函数最后执行即可
func init() {
	var err error
	path := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", IP, ":", PORT, ")/", dbName, "?charset=utf8"}, "")
	common.Db, err = sql.Open("mysql", path)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Context
	ctx := context.Background()

	// ping
	err = common.Db.PingContext(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Connected!")
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &middleware.BasicAuthMiddleware{},
	}
	// 静态资源 返回一个handler,它从请求URL中去掉指定的前缀,然后再调用另一个Handler
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./wwwroot"))))

	// 注册路由
	controller.RegisterRoutes()

	log.Println("Server starting...")
	// 性能分析
	go http.ListenAndServe(":8000", nil)

	// 此项目的 server
	server.ListenAndServe()
}
