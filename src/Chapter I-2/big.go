package main

import (
	"fmt"
	"math/big"
)

// int64 代表很大的数了 uint64 更大 但还是有数字超过这俩整数类型取值范围
// 我们用 float64 表示 ,如果没有为指数形式的数值指定类型 那么Go 将会将它视为float64类型

// 对于较大的整数 (超过10^18):big.int
// 对于任意精度的浮点类型:big.Float
// 对于分数 : big.Rat

// 题目: 俩种方式把 86000转化为 big.int 类型
func main() {
	// 方法一 是默认把int64位的转化成 big.Int 类型:
	// 缺陷: big.NewInt() 参数 不能超过 int64位取值范围,如果参数为 30e30  (30* 10^30)就会报错
	a := big.NewInt(86000)
	fmt.Println("big.NewInt 声明---", a)
	// 方法二 new(big.Int) 可以直接 参数为 30e30:30* 10^30
	b := new(big.Int)
	b.SetString("86000", 10) // 10 是10进制
	fmt.Println("new(big.Int)---", b)

	// 注意: Big.Int 一旦使用 等式内其他的部分也要使用big.Int
	result := new(big.Int)
	result.Div(a, b)
	fmt.Println("Div方法表示相处 俩边都要是 big.Int类型哦---", result)

}
