package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{"M", "N", "O", "P", "R"}

	// 一般 一个包里 大写字母的都默认导出
	// sort 为包 先调用 这个类型 StringSlice;
	// StringSlice() 转化为其类型 且 这个类型有一个Sort 方法可以直接调用
	sort.StringSlice(planets).Sort()
	fmt.Println(planets)
}
