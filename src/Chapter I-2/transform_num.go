package main

import (
	"fmt"
	"math"
)

// 类型不能混着用

// 字符串 和 数字 也不能混着用
func main() {
	/***数值类型间的转换 需要使用目标类型将其“包裹”起来***/

	// 整数 转 浮点型
	age := 18
	marAge := float64(age)
	fmt.Printf("Type %T \n", marAge)

	// 浮点转 整数 (可以从浮点类型转化为整数类型，小数点后边的部分会被截断，而不是舍入)
	earthDays := 365.2425
	fmt.Println("浮点转整数---", int(earthDays))

	// 整数 和浮点型 不能混着用 直接a+b 会报错
	a := 1
	b := -2.1
	fmt.Println("a+b", float64(a)+b)

	// 类型转换  环绕行为:
	var bh float64 = 32770
	var h = int16(bh) // 超过 int16最大值 32767 并且超过了3位
	// 结果计算 从最小值 + 超过的数-1  = -32768 + (3-1)
	fmt.Println("环绕行为 导致溢出---返回最小值", h)

	// 严谨的判断 通过 最大最小值来判断是否超过范围
	var isOverInt16 = bh < math.MinInt16 || bh > math.MaxInt16
	fmt.Println("是否超过 int16----", isOverInt16)
}
