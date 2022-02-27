package main

import "fmt"

// 指向结构体的指针
// 与字符串 和数值 不一样 复合字面量的前面可以放置 &
func main() {

	type person struct {
		name, superpower string
		age              int
	}
	// 表示: timmy指向的就是 复合字面值 person的 内存地址
	timmy := &person{
		name: "Timothty",
		age:  18,
	}

	// 访问字段的时, 对结构体进行解引用并不是必须的
	// 修改方式 不需要解引用
	timmy.superpower = "flying"
	// 解引用 方式2: (*timmy).superpower = "flying"

	fmt.Println("timmy---", timmy.age, timmy.superpower, timmy)

}
