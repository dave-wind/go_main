package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
1.JSON 与 Go 的 struct 相互转换
json:
{
	"id":123
	"name":"dave"
}
go struct:
type Person struct {
	ID     int
	NAme   string
}


2.映射 (字段导出需要大写!) Tags 方式:

type Person struct {
	ID    int     `json:id`
	Name  string   `json:name`
}
*/

/*
3.类型映射
Go bool: JSON boolean
Go float64 JSON 数值
Go string : JSON strings
Go nil: JSON null
*/

/*
4.未知结构 JSON:
map[string]interface{} 可以存任意类型JSON对象
[]interface{} 可以存储任意的 JSON数组
*/

/*
JSON操作 方法一:
5.读取JSON (一般读取request.body)

需要一个解码器: dec:=json.NewDecoder(r.Body) // 参数需要实现 Reader接口
在解码器上进行解码:dec.Decode(&query) // 把解码好的数据存放到query变量,这里传的是一个引用

6.写入JSON:(一般写到 response 返回)
需要一个编码器: enc:=json.NewEncoder(w) // 参数需要实现Writer接口
编码: enc.Encode(results)


7.JSON 与 go struct 转化方法二:
Marshal 和 Unmarshal

Marshal(编码): 把 go struct 转化为 json 格式
Marshallndent，带缩进
Unmarshal(解码): 把 json 转化为 go struct



8. 区别:
·针对 string 或 bytes:
· Marshal => String
. Unmarshal <= String

·针对 stream:
Decode <= Stream，从 io.Reader 读取数据
Encode => Stream，把数据写入到 io.Writer
*/

func main() {
	http.HandleFunc("/companies", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// 因为需要request的body 所以判断需要post方法
		case http.MethodPost:
			// 第一步 申明NewDecoder
			dec := json.NewDecoder(r.Body)
			// 申明一个model, 在model.go 文件里声明了 可以直接使用,前提是package 要一样
			company := Company{}

			// pointer struct 赋值给变量
			result := &company
			// 进行解码
			err := dec.Decode(result)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			// 现在拿到的是 go struct 所以还需要 编码为json response发出
			fmt.Println(result)
			// 修改request 内容
			result.Name = "Demo"
			// 编码
			enc := json.NewEncoder(w)
			err = enc.Encode(result)
			if err != nil {
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			// 默认不是post方法就返回错误
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe("localhost:8080", nil)
}

/*
操作:
1.go run .
2.安装REST Client 插件
3.在test.http 点击Send request
4.查看响应
*/
