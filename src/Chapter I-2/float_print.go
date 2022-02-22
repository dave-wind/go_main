package main

import (
	"fmt"
	"math"
)

// 精度问题 一般处理大型数据 需要用float32  math包里操作 全是 float64类型
func main() {
	fmt.Println("默认float 值 为 64:-----------:")
	// 默认 float64 占八个字节
	var pi64 = math.Pi
	// 单精度 占四个字节
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)

	fmt.Println("显示浮点数：默认尽可能展示多的显示几位小数-----------:")
	third := 1.0 / 3
	fmt.Println(third)

	fmt.Println("自定义显示几位小数 要用 Printf + f-----------:")
	fmt.Printf("%v\n", third)         //  这样没区别
	fmt.Printf("默认可能是6位:%f\n", third) // 默认可能是6位
	fmt.Printf(".3f自定义小数点后3位:%.3f\n", third)
	fmt.Printf("5.6f:显示最少字符个数为5位(包含小数点)若值长度不够加空格凑;小数点后有2位:%5.6f\n", third) // 长度5的
	fmt.Printf("06.2 显示最少字符个数6不够前面0来凑，小数点后面2位-----: %06.2f\n", third)
}
