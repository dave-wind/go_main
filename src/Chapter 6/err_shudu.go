package main

import (
	"errors"
	"fmt"
	"os"
)

const rows, columns = 9, 9

// Grid is a sudoku grid

type Grid [rows][columns]int8

// errors New 返回的是指针地址 而不是 字符串值
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// Set
func (g *Grid) Set(row, column int, digit int8) error {
	//  超过返回就抛出 error
	if !inBounds(row, column) {
		return ErrBounds
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

func main() {
	var g Grid
	err := g.Set(0, 0, 15)
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
