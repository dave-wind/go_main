package main

import (
	"fmt"
	"time"
)

/*
·Go 语言的函数和方法都是按值传递参数的，这意味着函数总是操作于被传递参数的副本。
当指针被传递到函数时，函数将接收传入的内存地址的副本。之后函数可以通过解引用内存地址来修改指针指向的值。
!!!!!!按值传递
*/

type person struct {
	name, superpower string
	age              int
}

// 声明一个 指向person的一个指针类型
func birthday(p *person) {
	// 传递进去 会声明一个 指针的副本 (但指向的内存地址都是一样的)
	// 这个指针副本 隐含了 先 解引用 再 age++
	p.age++
}

func birthdayNotPointer(p person) {
	p.age += 1
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}
	// 值传递 修改都互不影响
	rebecca2 := rebecca

	birthday(&rebecca)

	// 传递的是 复合字面体;非指针 p.age 并没有加加
	birthdayNotPointer(rebecca2)
	// rebecca2.age += 1

	fmt.Printf("%+v\n", rebecca)
	// 并没有变
	fmt.Println(rebecca2)

	fmt.Println(" 指针接受者----------------------")

	terry := &per{
		name: "Terry",
		age:  15,
	}

	terry.birthday()
	fmt.Printf("%+v\n", terry)

	nathan := per{
		name: "Nathan",
		age:  17,
	}
	// !!!!! 点标记法调用 方法 就会自动 使用&符号 来取变量内存地址:
	//  ==》(&nathan).birthday()
	nathan.birthday()

	fmt.Printf("%+v\n", nathan)

	fmt.Println("time---------")
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	tommorrow := day.Add(24 * time.Hour)
	fmt.Println("当前日期", day.Format(layout))
	fmt.Println("明天日期", tommorrow.Format(layout))

}

// 指针接受者:
// 方法的接收者 和 方法的参数 在处理指针方法是很相似的
type per struct {
	name string
	age  int
}

// 指向per的一个指针类型,作为方法的接收者
// go语言里 往方法和函数里 通常以指针来传递;且必须定义指针类型;才可以对内存地址里面的key进行直接操作
func (p *per) birthday() {
	p.age++
}

/*
注意
使用指针作为接收者的策略应该始终如一:
如果一种类型的某些方法需要用到指针作为接收者，就应该为这种类型的所有方法都是用指针作为接收者。
*/
