package main

import "fmt"

// 指针是指向 [另一个变量地址]的变量
// Go语言的指针同时强调安全性,不会出现迷途指针 (dangling pointers)

//RAM: 指随机存取存储器（random access memory，RAM）又称作“随机存储器”，是与CPU直接交换数据的内部存储器，也叫主存(内存)
// 每声明一个变量,它们就会将其值存储在计算机的RAM中,存储的位置 就是该变量的内存地址
// & 表示地址操作符,通过 & 可以获得变量的内存地址

// 注: &操作符 无法获得字符串/数值/布尔字面值的地址 直接 &42 &"dave" 都会编译报错

// *操作符 与 & 作用相反, 它用来解引用,提供 内存地址指向的值
// Go 语言 不允许 address++ 这样的指针运算 进行操作
func main() {
	// demo := []string{"deno"}
	answer := 42
	fmt.Println("& 操作符:", &answer)

	address := &answer
	fmt.Println(" * address的解引用", *address)
	fmt.Println(" *&answer 是其本身", *&answer)
	/*
		// address 是 answer变量的内存地址, 通过 *address 解引 为42
		address = 0xc0000b2008
		                     |
		answer = 	42
	*/

}
