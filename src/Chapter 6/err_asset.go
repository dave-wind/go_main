package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type SudokuError []error

func (se SudokuError) Error() string {
	var s []string
	for _, err := range se {
		s = append(s, err.Error())
	}
	return strings.Join(s, ",")
}

const rows, columns = 9, 9

type Grid [rows][columns]int8

var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

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

// 类型断言 可以吧接口类型转化成底层的具体类型, 使用断言类型来访问每一种错误
// 如果类型满足多个接口,那么类型断言可使它从一个接口类型转化为另一种接口类型
func main() {
	var g Grid
	err := g.Set(12, 0, 15)
	if err != nil {
		// err 如果是 SudokuError,ok为true
		if errs, ok := err.(SudokuError); ok {
			fmt.Printf("%d error(s) occurred:\n", len(errs))
			for _, e := range errs {
				fmt.Printf("- %v\n", e)
			}
		}
		os.Exit(1)
	}

}
