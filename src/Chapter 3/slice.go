package main

import (
	"fmt"
	"strings"
)

// 假设planets 是一个数组，那么 planets[0:4] 就是一个切片 它切分出了数组的前4个元素
// 切片数组不会导致数组被修改 它只是创建了指向数组的 一个窗口 或视图 这种视图就是 切片
func main() {
	planets := [...]string{"a", "b", "c", "d", "e", "f"}

	// slice 使用的是半开区间 包含 0 不包含 4
	slice := planets[0:4] //
	fmt.Println(planets, slice)

	// slice的默认索引：
	// 忽略掉 slice的起始索引 表示从数组的起始位置进行切分，忽略掉结束索引 反之亦然
	fmt.Println("忽略掉索引：", planets[:4], planets[1:])

	// 如果同时忽略掉 起始和结束的索引 那就是包含数组所有元素的一个slice
	fmt.Println("同时忽略掉 起始和结束的索引", planets[:])

	// 注: slice 的索引 不能为负数

	/*************************slice 切分字符串***************************************/

	neptune := "Neptune"
	tune := neptune[3:]
	fmt.Println("字符串切片:tune", tune)

	// 切分字符串 索引代表的是字节数 而非rune的数!!!! 一般只有用 库:utf8.RuneCountInString 才返回真实的rune数
	question := "emésydivirtete"
	fmt.Println("字符串切片索引切的是字节:", question[:6]) // emésy

	// 直接声明 slice : []string 因为 slice 长度是可变的
	simple := []string{"Venus", "Earth"}
	fmt.Printf("%T %T \n", planets[:], simple)
	fmt.Printf("数组类型:  %T \n", planets)

	// 合并字符串:
	pla := []string{" Venus   ", "Earth   ", " Mars"}
	hyp(pla)
	fmt.Println("join 合并:", strings.Join(pla, "-"))

}

func hyp(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
