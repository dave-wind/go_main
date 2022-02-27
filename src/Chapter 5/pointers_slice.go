package main

import "fmt"

// slice 指向数组
/*
·之前说过 slice 是指向数组的窗口，实际上slice 在指向数组元素的时候也使用了指针。
。每个 slice 内部都会被表示为一个包含3个元素的结构，它们分别指向:
数组的指针 slice 的容量 slice 的长度
·当 slice 被直接传递至函数或方法时，slice 的内部指针就可以对底层数据进行修改。
*/

// 指向slice的显式指针的唯一作用 就是修改 slice本身,slice长度,容量 以及起始偏移量

// 参数: 指向切片字符串的一个指针类型
func reclassify(planets *[]string) {
	// 解引用操作 进行修改
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune", "Pluto",
	}
	reclassify(&planets)
	fmt.Println(planets)
}

// 注: 如果函数和方法想修改他们接受的数据,那么他们应该使用 切片指针 (可修改!!!)
