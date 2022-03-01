package main

import "fmt"

/*
!!!保护你的方法
* 避免 nil 引发 panic !!
*/

/*
*因为值为 nil 的接收者和值为 nil 的参数在行为上并没有区别
，所以 Go 语言即使在接收者为 nil 的情况下，也会继续调用方法。
*/

type person struct {
	age int
}

// 指针为nil做接受者; 也可以作为nil做函数参数
func (p *person) birthday() {
	// 如果函数内部不容错 就会 运行时 panic;所以需要容错
	if p == nil {
		return
	}
	// 这里会报 panic,空指针 不可 取属性值
	fmt.Println("p.age", p.age)
	p.age++
}
func main() {
	// 类型为指向person的类型, 值为nil
	var nobody *person
	fmt.Printf("nobody %T,%v\n", nobody, nobody)

	//可以调用
	nobody.birthday()

}
