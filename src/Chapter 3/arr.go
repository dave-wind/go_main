package main

import (
	"fmt"
)

func main() {
	// 数组是一种固定长度且有序的元素集合
	var planets [8]string

	fmt.Println("数组声明 固定长度：", planets)

	// 访问数组元素
	planets[0] = "a"
	planets[1] = "b"
	planets[2] = "c"
	fmt.Println("planets length 用 len 访问长度", len(planets))

	// 声明数组 未被赋值的元素值就是对应类型的零值：
	fmt.Println("字符串 数组没赋值的为空", planets[7] == "")
	var un [8]uint8
	fmt.Println("数字 数组没赋值的为0", un[7] == 0)

	// 访问一个数组内的值
	earth := planets[2]
	fmt.Println("earth", earth)

	// 数组越界 越界 编译的时候就会报错，如果编译不报错 panic：运行时候也会报错
	// planets [8] = "d"

	// 复合字面值初始化数组
	dwarfs := [5]string{"a", "b", "c", "d", "e"}
	dwarfs2 := [...]string{"a", "b", "c"} // 前方 ... go编译器会自动算出元素数量
	fmt.Println("复合字面值初始化数组:", dwarfs, dwarfs2)

	// 遍历数组
	for i := 0; i < len(dwarfs); i++ {
		fmt.Println("for循环遍历", i, dwarfs[i])
	}

	for _, item := range dwarfs2 {
		fmt.Println("range循环遍历", item)
	}

	// 与 for 不同 range 对每个迭代值都创建了一个拷贝 数据大的话 性能会没有 for 好
	for i := range dwarfs2 {
		fmt.Println("range循环遍历 只是遍历索引 每次迭代元素 会占用内存", dwarfs2[i])
	}

	// 数组的copy
	// 先copy
	copyArr := dwarfs2
	// 再修改原数组
	dwarfs2[1] = "修改"
	// 结果并不会 改动copy过的数组
	fmt.Println("数组的copy", copyArr, dwarfs2)

	// 数组作为值 无论赋值给新的变量 还是 当参数 传递给函数；都会产生一个完整的数组副本
	// 数组当作参数 传递给函数 效率非常低效：（因为都要产生一个副本）
	// 所以 不会用 数组 作为参数传递给函数 用切片！！！slice

	// 多维数组

	var board [3][4]byte

	board[0][0] = 'r' // 第一个数组 item 里面 第一个元素 赋值
	board[0][3] = 'r' // 第一个数组 item 里面 最后一个元素 赋值
	fmt.Println("board 多维数组：", board)

	for i := range board[1] {
		board[1][i] += 'A'
	}
	fmt.Println("多维数组 循环遍历第二个元素 每个加p：", board)
}
