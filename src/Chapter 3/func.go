package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// rand 包的Intn: func Intn(n int) int 用法:
	num := rand.Intn(10)
	fmt.Println("num--", num)

	// 注意:
	// 在Go里 大写字母开头的函数!!!! 变量或其他的标识符都会被到除对其他包可用,小写不行

	// 函数参数可多个: 形参类型相同 可以只写一次
	// func Unix (sec int64 nsec int64) Time ==> func Unix (sec, nsec int64) Time

	// 函数返回值可多个: func Atoi(s string)(i int, err error)
	// 函数返回值 可直接只写类型:  func Atoi(s string)(int, error)
	contdown, err := strconv.Atoi("10") // (字符串 转 数字)
	fmt.Println(contdown, err)

	// 函数声明 可变参数函数: ...interce{} :空接口(任何类型)  ... :代表参数数量可变
	// func Println(a ...interce{})(n int err error)

}
