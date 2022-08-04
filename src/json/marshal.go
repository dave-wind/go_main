package main

import (
	"encoding/json"
	"fmt"
)

/*
Marshal(编码): 把 go struct 转化为 json 格式
Marshallndent，带缩进
Unmarshal(解码): 把 json 转化为 go struct

*/

type Company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

func main() {
	jsonStr := `
	{
		"id": 123,
		"name":"Google",
		"country":"USA"
	}
	`
	c := Company{}

	// 解码, 第一个参数:byte slice
	_ = json.Unmarshal([]byte(jsonStr), &c)
	fmt.Println(c)

	// go struct 转 json
	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	// go struct 转 json,且带空格, 第二个参数:前缀,第三个:缩进
	bytes1, _ := json.MarshalIndent(c, "", " ")
	fmt.Println(string(bytes1))
}
