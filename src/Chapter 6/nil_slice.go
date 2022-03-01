package main

import "fmt"

/*
* 如果 slice 在声明之后没有使用复合字面值或内置的 make 函数进行初始化，那么它的值就是 nil。
·幸运的是，range、len、append 等内置函数都可以正常处理值为 nil的 slice。
*/

func main() {
	// 默认值为 nil
	var soup []string

	fmt.Println(soup == nil) // true

	// nil 不会走range 也不会报错的
	for _, ingredient := range soup {
		fmt.Println(ingredient)
	}

	// 也不会报错 输出为 0
	fmt.Println(len(soup))

	// nil 直接append 就累加值了
	soup = append(soup, "onion", "carrot", "celery")
	fmt.Println(soup)

	fmt.Println("空slice 和值为nil 的slice---------------:")

	sou := mirepoix(nil)
	fmt.Println(sou)
}

// 虽然空slice 和值为nil 的slice 并不相等,但它们通常可以替换使用
func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}
