package main

import "fmt"

// 结构类型 struct

// Go 允许你将不同类型的东西组合在一起 类似: 类
// 可以声明变量 也可以声明 类型; 类型的好处是 可以复用
// type location struct {
// 	lat, long float64
// }

type location struct {
	lat  float64
	long float64
}

// 别名
type dis location

func distance(loc1, loc2 location) dis {
	return dis{0.0, 0.0} //初始化赋值

}
func main() {
	a := location{lat: 3.3}
	b := location{long: 6.6}
	fmt.Println("distance(a, b)", distance(a, b))

	// 符合字面值初始化
	curiosity := location{-4.5895, 137.4417}
	fmt.Println("第一种声明: 按字段定义的顺序进行初始化:", curiosity)

	opportunity := location{lat: -4.5895, long: 137.4417}
	fmt.Println("第二种声明: 成对的字段和值进行初始化:", opportunity)

}
