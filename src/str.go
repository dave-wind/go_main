package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {

	str := "peace"

	// 字符串 字面值 可以包含 转义字符 例如 \n
	fmt.Println("peace be upon you \nupon you be peace", str)
	//  如果非要 输出 \n  用 ``
	fmt.Println(`str----------- \n`, str)

	// 用 `` 这种方式 可以实现换行
	fmt.Println(`peace be upon you
nupon you be peace`)

	// 类型别名 就是同一个类型另一个名字
	// 如下为uint8 和 int32 类型的别名 语法如下:
	// type byte = uint8
	// type rune = int32

	// %c 将会打印 的是字符!!: rune 表示的话 那就代表他是一个字符 用单引号 括起来
	var pi rune = 960
	var alpha rune = 940
	var omg rune = 963
	var bangd byte = 33
	fmt.Printf("展示字符: %c %c %c %c \n", pi, alpha, omg, bangd)

	/****************字符**********************/
	// 如果没指定字符类型的话，那么Go 会推断它的类型为rune
	grade := 'A'
	fmt.Println("grade 这里的 grade仍然包含一个数值，本例中就是65，它是A的code point。", grade)
	var star byte = '*'
	fmt.Println("42是 *的code point", star)
	var smile = '😊'
	fmt.Println("128522是 😊的code point", smile)
	var e = 'é'
	fmt.Println("233是 é的code point", e)

	// 字符串值不能改
	msg := "shalom"
	cc := msg[5]
	fmt.Printf("字符串 message[5] 可用 c打印: %c\n", cc)
	// mesag[5] = 'd' // 不可修改

	// Caesar cipher 凯撒加密法
	c := 'a'
	c = c + 3 // ASCII code point 一直往后累加
	fmt.Printf("c+3 =:%c\n", c)
	if c > 'z' { //超过 z 回到 A
		c = c - 26
	}
	fmt.Printf("after c:%c\n", c)

	// 字符之间 加减 demo:
	demo := 'g' - 'a' + 'A' // 相当于 code point 进行操作
	// demo := '!' + 'A'
	fmt.Printf("g-a+A = %c\n", demo)

	message := "shorm"
	fmt.Println("len 内置函数 无需 import len返回的是变量的byte数", len(message), unsafe.Sizeof(message))

	/****************RuneCountInString && DecodeRuneInString**********************/
	/*
		本例中，程序访问的是message这个字符串的每个字节（8位），可能没有考虑多字节的情况（16、32位）。•
		如何支持西班牙语、俄语、汉语等？•把字符解码成rune类型，然后再进行操作。•
		使用utf-8包，它提供可以按rune计算字符串长度的方法。•
		DecodeRuneInString 函数会返回第一个字符，以及字符所占的字节数。•
		所以Go 里的函数可以返回多个值。
	*/
	// 对于 非英文字符串 可以用 RuneCountInString 返回 真实字节长度
	fmt.Println(utf8.RuneCountInString(message), "bytes")
	r, size := utf8.DecodeRuneInString(message)
	fmt.Println("message", r, size)

	/****************range****************/
	question := "tíaíing"
	for i, q := range question {
		// i 是索引 ,q 是字符值  这里 i 可以用 丢弃标识符 _ 只输出 字符值
		fmt.Printf("%v %c\n", i, q)
	}
	fmt.Println("西班牙语 tíaíing--- 纯字符串字符长度:", utf8.RuneCountInString(question), "bytes")
	// 因为有的 字符 是按 16位 或者 32位来的
	fmt.Println("西班牙语 tíaíing--- 字节长度:", len(question))

	// 题目 把如下字符串 每个字母加3 会变成什么结果:
	ti := "Afdph, Lvdz"
	var ti_change string
	for _, t := range ti {
		t = t + 3
		if t > 'z' {
			t = t - 26
		}
		ti_change += string(t)
	}
	fmt.Println("字符串累加结果:---", ti_change)
	// fmt.Printf("字符串累加结果 %c\n", ti_change)
}

// 字符，code points, runes, bytes
// .Unicode 联盟为超过 100 万个字符分配了相应的数值，这个数叫做 code point.
// ·例如:65 代表 A，128515代表
// 。为了表示这样的 unicode code point，Go 语言提供了 rune 这个类型:它是 int32 的一个类型别名。
// ·而 byte 是 uint8 类型的别名，目的是用于二进制数据。
// ·byte 倒是可以表示由 ASCII 定义的英语字符，它是 Unicode 的一个子集(共128个字符)

/*

Go中的字符串是用UTF-8编码的，UTF-8是Unicode Code Point的几种编码之一。•
UTF-8是一种有效率的可变长度的编码，每个code point可以是8位、16位或32位的。•
通过使用可变长度编码，UTF-8使得从ASCII的转换变得简单明了，因为ASCII字符与其UTF-8编码的对应字符是相同的。•
UTF-8是万维网的主要字符编码。
它是由Ken Thompson于1992年发明的，他是Go语言的设计者之一。

*/
