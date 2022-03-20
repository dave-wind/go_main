package main

import "fmt"

/*
@ D & C 策略
1.找到一个简单的基线条件 (Base Case)
2.把问题分开处理,直到它变为基线条件
*/
// 基线条件: 空数组[],其和为 0
/* 递归: [1,3,5]
1 + SUM([3,5])
  3 + SUM([5])
    5 + SUM([])
*/

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	// 切片,先取出第一个 + 递归()
	return arr[0] + sum(arr[1:])
}

func main() {

	total := sum([]int{1, 3, 5, 7, 9})
	fmt.Println("求和", total)
}
