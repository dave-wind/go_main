package main

import "fmt"

func main() {
	fmt.Println("123") // Println 可换行
	fmt.Println("456")
	fmt.Print("this is Print 1 ") // Print不换行
	fmt.Print("this is Print 2 \n")

	fmt.Print("xxx", 123, 456, "\n")

	fmt.Println("--------------------动态穿参数Log：")

	// Printf 第一个参数必须是字符串 后面 一个个的传参
	fmt.Printf("%v 哈哈 %v \n", "dave jones", "ddd")

	fmt.Println("--------------------加空格 10v：代表变量前面加空格 -10v：代表变量后面加空格")

	// -10 代表 右边 多 10个空格  10:代表 左边 多10个空格
	fmt.Printf("%-10v && %10v \n", "Dave", "Jones")

	fmt.Printf("%10v && %-10v \n", "Dave", "Jones")

	fmt.Println("-------------------- 变量几种声明方式:")

	var a = 10
	var b = 12
	fmt.Println(a, b)

	var c, d = "c", "d"
	fmt.Println(c, d)

	var (
		e = "e"
		f = "f"
	)
	fmt.Println(e, f)

	fmt.Println("-------------------- go里没有 ++age这样的写法,会报错 error: expected ++")
}
