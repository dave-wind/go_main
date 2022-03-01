package main

import (
	"fmt"
)

/*
 ·Nil 是一个名词，表示“无”或者“零”。·在Go 里，nil 是一个零值。
如果一个指针没有明确的指向，那么它的值就是nil·除了指针，nil 还是 slice、map 和接口的零值。
Go 语言的 nil，比以往语言中的 null 更为友好，并且用的没那么频繁，但是仍需谨慎使用。
*/
func main() {
	// 声明一个 指针 类型为int, 值为 nil
	var nowhere *int
	fmt.Println(nowhere)

	// 尝试解引用一个nil指针 将会导致 程序崩溃 (panic)这里就会运行时报错:
	// fmt.Println(*nowhere)

	// 所以需要判断 容错

	if nowhere != nil {
		fmt.Println(*nowhere)
	}
	var str *string
	fmt.Println("str *string 也为:", str)

}
