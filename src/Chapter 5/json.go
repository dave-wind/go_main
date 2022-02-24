package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Println("error--", err)
		// 不等于 0 代表发生了故障
		os.Exit(1)
	}

}
func main() {
	// 大写了字段,表示字段是可以导出的,小写不能导出
	type location struct {
		Lat, Long float64
	}
	curiosity := location{Lat: -4.589, Long: 3.758}

	// json.Marshal 可以把 curiosity 转化为 bytes; 前提是 struct 字段是能导出的 要大写!
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(bytes)

	// 再 通过 字符串化 转为 JSON
	fmt.Println(string(bytes))

	// 使用 struct标签 来自定义JSON
	// Go 语言中的 ison 包要求 struct 中的字段必须以大写字母开头，类似 CamelCase 驼峰型命名规范。
	// 但有时候我们需要 snake case!!! 蛇形命名规范，那么该怎么办? 可以为字段注明标签，使得 json 包在进行编码的时候能够按照标签里的样式修改字段名。
	type location2 struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	// 只要跟换 类型 即可 输出:
	fmt.Println(`{"latitude":-4.589,"longitude":3.758}`)

	// 问题? go 编码json数据 为什么非要字段大写
}
