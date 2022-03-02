package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// 自定义错误类型： error类型是一个内置接口，任何类型只要实现了 返回string的 Error() 方法就满足了该接口
// 可以创建新的错误类型,
// 按照惯例： 自定义错误类型 名称应以 Error结尾
type SudokuError []error // error slice 类型(集合)

// 自定义错误类型的 Error 方法

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

const rows, columns = 9, 9

// Grid is a sudoku grid

type Grid [rows][columns]int8

// errors New 返回的是指针地址 而不是 字符串值
// 按照惯例：自定义错误变量 要以 Err开头
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// Set
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError

	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}

	if !validateDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}

	g[row][column] = digit

	return nil
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func validateDigit(n int8) bool {
	if n < 9 {
		return true
	}
	return false
}

func main() {
	var g Grid
	err := g.Set(12, 0, 15)
	if err != nil {
		switch err {
		case ErrBounds, ErrDigit: // 判断的是指针地址
			fmt.Println("Les erreurs de paramètres hors limites.")
		default:
			fmt.Println(err)
		}
		os.Exit(1)
	}

}
