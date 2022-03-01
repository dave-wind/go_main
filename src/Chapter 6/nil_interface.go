package main

import "fmt"

/*
* 。声明为接口类型的变量在未被赋值时，它的零值是 nil。
对于一个未被赋值的接口变量来说，它的接口类型和值都是nil，并且变量本身也等于 nil。
。当接口类型的变量被赋值后，接口就会在内部指向该变量的类型和值。
*/

// type person struct {
// 	age int
// }

type demoer interface {
	String()
}

func main() {
	// 声明一个空接口 类型为 nil 值为nil, 当值和类型都为nil 才等于 nil
	var v interface{}
	fmt.Printf("%T %v %v\n", v, v, v == nil)
	// 声明一个 指针
	var p *int
	// 赋值给v ;v会发生变化, 类型就会 *int(指向int类型的指针)
	v = p
	fmt.Printf("%T %v %v\n", v, v, v == nil)

	// 检验接口变量的内部表示
	// %#v 可以表示内部显示: (*int)(nil)
	fmt.Printf("%#v\n", v)

	// 定义一个接口,声明变量为 接口 值为nil 类型也是nil
	fmt.Println("demo------------:")
	var (
		s demoer
		f fmt.Stringer
	)
	fmt.Printf("%T,%v\n", s, s)
	fmt.Printf("%T,%v\n", f, f)
}
