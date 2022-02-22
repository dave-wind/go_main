package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 生成随机日期

var erase = "AD:---------"

func randNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func main() {
	year := randNum(1979, 2042)
	fmt.Println(erase, year)
	daysInMonth := 31
	// month :
	// fmt.Println(year)
	month := randNum(1, 12)
	switch month {
	case 2:
		if year%100 != 0 && year%4 == 0 || year%400 == 0 {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	fmt.Println(erase, year, month, daysInMonth)

}
