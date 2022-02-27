package main

import (
	"fmt"
)

// 指向 数组的指针

/*
和结构体一样，可以把 & 放在数组的复合字面值前面来创建指向数组的指针。
数组在执行索引或切片操作时会自动解引用。没有必要写(*superpower)[0] 这种形式。
·与C语言不一样，Go里面数组和指针式两种完全独立的类型。
Slice 和 map 的复合字面值前面也可以放置&操作符，但是 GO 并没有为它们提供自动解引用的功能。
*/
func main() {
	superpowers := &[3]string{
		"fight", "invisibility", "super strength",
	}
	fmt.Println(superpowers)
	fmt.Println(superpowers[0])
	fmt.Println(superpowers[1:2])

	// 修改数组
	var board [8][8]rune
	reset(&board)
	// fmt.Println(len(board), cap(board))
	fmt.Printf("%c", board[0][0])
}

func reset(board *[8][8]rune) {
	board[0][0] = 'r'
}
