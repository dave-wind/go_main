package main

import "fmt"

// 指针类型用处:
// 用在所有需要用到类型的地方，如变量声明、函数形参、返回值类型、结构体字段等。
func main() {
	answer := 42
	var address *int
	// 指针存储的是内存地址
	address = &answer // *int 代表的是: 指向int类型指针类型 (*代表的是指向)
	fmt.Printf("address is a %T\n", address)

	canada := "Canada"

	// 1. 将*放在类型前面 表示声明指针类型
	// 2.将*放在变量前面表示解引操作

	var home *string // 指向字符串类型的指针类型
	fmt.Printf("home is a %T\n", home)

	// !!!! 通俗解释: home 就是指向 变量canada的内存地址, canada 改变 *home也会改变
	home = &canada                // 指针赋值
	fmt.Println("解引 home", *home) // 解引

	// 指针就是用来指向的, 不是指向副本 俩种写法:
	// 1.修改变量
	canada = "Change Canada"
	fmt.Println("修改后的指针:", *home)
	// 2.直接修改 *home
	*home = "直接修改 *home"
	fmt.Println("修改*变量:", canada)

	major := home // 指针赋值给 另一个变量
	fmt.Println("指针赋值给 另一个变量 major:", major)

	*major = "This is Major"
	fmt.Println("修改另一个指针 *major 看初始变量也会改变:", canada)

	// 判断指针是否相等 看 指针指向的内存地址是否相等
	// major 和 home 都是 指向 canada的内存地址 所以一样
	fmt.Println("major == home:", major == home)

	lightfoot := "Robert M Lightfoot Jr."
	home = &lightfoot
	// 因为 home指针 现在指向了 lightfoot,而 major 指向的还是以前的 This is Major
	fmt.Println("major == home:", major == home)

	// 只传递副本 修改指针不影响
	charles := *major
	*major = "Charles Bolden"
	fmt.Println(charles)
	// 这里 canada 改变 是因为 home指针 指向canada,虽然这里改变的是major
	// 但是 major := home 指针赋值了 所以产生了联系 俩指针指向同一个地址, 所以 major修改内存地址. canada就会变
	fmt.Println("canada", canada)

	charles = "Charles Bolden"
	fmt.Println(charles == canada)   // 值可以相等
	fmt.Println(&charles == &canada) // 指针不相等
}
