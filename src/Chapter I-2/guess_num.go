package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	var result int = rand.Intn(max-min+1) + min
	return result
}

func main() {
	var guessNum int = 15
	var randomNum int = randNum(1, 100)

	// for times := 0; randomNum != guessNum; times++ {
	// 	if guessNum > randomNum {
	// 		fmt.Println("比猜的数小了--- 随机数是：", randomNum)
	// 	} else {
	// 		fmt.Println("比猜的数大了--- 随机数是：", randomNum)
	// 	}

	// 	randomNum = randNum(1, 100)
	// }
	var times = 0
	for randomNum != guessNum {
		if guessNum > randomNum {
			fmt.Println("比猜的数小了--- 随机数是：", randomNum)
		} else {
			fmt.Println("比猜的数大了--- 随机数是：", randomNum)
		}
		times++
		randomNum = randNum(1, 100)
	}
	fmt.Printf("Go版 结果---你一共用了 %v 次猜对了是 %v \n", times, randomNum)
}
