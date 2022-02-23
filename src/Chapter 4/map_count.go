package main

import "fmt"

func main() {
	// 题目: 把 temperatures 里面出现的度数 进行计数
	//定义一个切片
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// 对象进行预分配
	frequency := make(map[float64]int)

	// 	特性: 把切片的value 作为key 当对象没有这个key 初始化为0
	// 当有这个key 就会 ++ 累加
	for _, t := range temperatures {
		frequency[t]++
	}

	// range 遍历 对象 顺序是不可保证的
	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times \n", t, num)
	}
	// %d: 输出十进制整数
	// %f: 输出十进制浮点数
	// %s: 输出字符串
	// %x: 输出十六进制
}
