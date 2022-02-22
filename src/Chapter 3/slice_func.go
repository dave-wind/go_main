package main

import (
	"fmt"
)

// 在Go 里 可以将 slice 或 数组 作为底层类型,然后绑定其他方法

// 自定义一个 类型 底层类型为 切片
type CustomStringsSlice []string

type CustomInt []int

// 带有方法的切片
func (k CustomStringsSlice) Sort(s []string) []string {
	for i := range s {
		// item := k[i]
		s[i] += k[0]
	}
	return s
}

func main() {

	planets := []string{"M", "N", "O", "P", "R"}
	fmt.Println("planets[0:4]:", planets[0:4])

	// slice 类型
	var d = CustomStringsSlice{"-I"}
	fmt.Println("用变量直接调用方法", d.Sort(planets[:]))
	fmt.Println("类型调用方法:", CustomStringsSlice(planets[0:1]).Sort(planets[:]))

	var a = []int{1}
	// 括号
	fmt.Println("CustomInt []int:", CustomInt(a))

}
