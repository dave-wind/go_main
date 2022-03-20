package main

import "fmt"

/*
* Quicksort 排序数组
1. 选择一个 pivot 基准值
2. 数组分区:
	左侧数组元素 都比 pivot 小
	右侧数组元素 都比 pivot 大
3. 在俩个子数组上递归调用 quicksort
*/

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	// var index = int(len(arr) / 2)
	pivot := arr[0]
	var left, right []int
	for _, ele := range arr[1:] {
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	// left + 中间 + right ...为展开操作符
	return append(quicksort(left),
		append([]int{pivot}, quicksort(right)...)...)
}

func main() {
	arr := []int{9, 3, 10, 8, 4, 25, 33, 97, 42}
	result := quicksort(arr)
	fmt.Println(result)
}
