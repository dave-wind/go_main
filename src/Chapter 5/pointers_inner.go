package main

import "fmt"

/*
·Go 语言提供了 内部指针 这种特性。
·它用于确定结构体中指定字段的内存地址。。(例子 26.15)
·& 操作符不仅可以获得结构体的内存地址，还可以获得结构体中指定字段的内存地址。
*/

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

func main() {
	player := character{name: "Matthias"}
	// 提供了结构体内部指定字段的指针!!!!!
	levelUp(&player.stats)
	fmt.Printf("%+v\n", player.stats)

	fmt.Println("隐式指针----------------------")

	/*
			隐式的指针
		·Go 语言里一些内置的集合类型就在暗中使用指针。
		·map 在被赋值或者呗作为参数传递的时候不会被复制。
		map 就是一种隐式指针。
		。这种写法就是多此一举:func demolish(planets *map[string]string) map 的键和值都可以是指针类型
		需要将指针指向 map 的情况并不多见
	*/

}
