package main

import (
	"fmt"
	"strings"
)

// 接口:
// 1.接口关注于类型可以做什么,而不是存储什么!!
// 2.接口通过列举类型必须满足一组方法来进行声明
// 3.在Go语言中 不需要显式声明接口,都是隐式满足的

// 定义一个接口变量
var t interface {
	talk() string
}

type martian struct{}

// 只要任何类型满足了 接口的要求 就会成为 接口的值
// 表示 martian类型  满足了 || 实现了 这个接口
func (m martian) talk() string {
	return "nack nack"
}

// laser 满足了 || 实现了 这个接口
type laser int

func (l laser) talk() string {
	return strings.Repeat("pew:", int(l))
}

func main() {
	// 类似java的多态 因为 t 可变为不同类型
	t = martian{}
	fmt.Println(t.talk())

	// 3 int类型 转成了 laser类型 然后赋值给 t; t再调用方法
	t = laser(3)
	fmt.Println(t.talk())

}
