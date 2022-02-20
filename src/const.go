package main

import "fmt"

//
func main() {
	// 如果没有为指数形式的数值指定类型 那么Go 将会将它视为float64类型
	const distance0 = 24e30
	fmt.Printf("常量 distance----数据类型 %T for distance \n", distance0) // float64

	// 常量非科学计数法
	// 常量 本身为🈚类型(整数) untyped int; 但是作为整数默认就会是int类型 所以用 Println 执行就会认为int类型就会报错
	const distance = 24000000000000000000
	// fmt.Println("非科学计数法 distance", distance) //过大报错: constant 24000000000000000000 overflows int

	// 比较大的常量 可以直接使用(作为字面值)
	// 针对字面值和常量的计算是在编译阶段完成的
	fmt.Println("执行--", distance/297777/4399)

	//作业 大矮星是已知距离地球最近星系 距离太阳:
	total_distance := 236000000000000000 // 公里
	year_distance := 300000              // 一光年公里
	result := total_distance / year_distance
	fmt.Println("此距离大概有----", result, "光年")
}

// 对于较大的数值常量 尽管Go编译器使用Big包来处理无类型的常量, 但是常量与big类型是不能互换的
