package main

import (
	"fmt"
	"strconv"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	// 	输出的 是 code point 对应的字符
	fmt.Println(string(pi), string(alpha), string(omega), string(bang))

	/******* 数值 转 字符串 ***********/

	// 想把数值转化为 string  它的值必须能转化为  code point
	var A = 65
	fmt.Println("65 转为字符串", string(A))

	// strconv 包 的 Itoa 函数就可以做这件事:
	countDown := 10
	// strconv == string convert 字符串转换
	str := "Start  " + strconv.Itoa(countDown) + "  end"
	fmt.Println("结果", str)
	// Itoa 是  Integer to ASCII  的意思
	//  Unicode  是 ASCII的超集 他们前 128个 code points 是一样的 (数字 英文字母 常用标点)

	//  另外一种把数值转化为 string 为Sprintf 和 Printf 略类似 但是!!会返回一个 String
	countDown2 := 10
	str2 := fmt.Sprintf("Sprintf %v end", countDown2)
	fmt.Println("str2---", str2)

	/***** 字符串 转 数字 *****/
	// 首先 他转化后 要能是 数字 有点像node 😂 err
	strdemo := "123"
	str2num, err := strconv.Atoi(strdemo)
	if err != nil {
		fmt.Println(err.Error())
	}
	// err == nil 代表没有错
	fmt.Println("strconv.Atoi---", str2num, "错误:", err)

	// Go 是静态类型语言，一旦某个变量被声明，那么它的类型就无法再改变了

	/******* 布尔值进行转化 *******/
	//  注意：如果你想使用  string(false), int(false)  bool(1), bool("yes")
	//  等类似的方式进行转换，那么 Go 编译器会报错
	//  某些语言里，经常把 1 和 0 当作 true 和 false 但是Go 语言不行

	//  写一个程序，把布尔类型转化为整数类型，true->1 ,false->0
	bol := fmt.Sprintf("%v", false)
	var bol_str byte
	if bol == "true" {
		bol_str = 1
	} else {
		bol_str = 0
	}
	fmt.Println("bol_str:", bol_str)

	//  写一个程序，把字符串转化为布尔类型： “true”, “yes”, “1” 是 true  “false”, “no”, “0” 是 false 针对其他的显示错误信息
	arr1 := []string{"false", "0", "no"}
	arr2 := []string{"true", "1", "yes"}
	guess := "1"
	var guess_result bool
	if contains(arr1, guess) {
		guess_result = false
	} else if contains(arr2, guess) {
		guess_result = true
	} else {
		fmt.Println(fmt.Errorf("错误值", guess_result))
		return
	}
	fmt.Printf("guess_result-- %v ,数据类型: %T\n", guess_result, guess_result)
}
