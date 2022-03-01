package main

import "fmt"

// 和slice 一样 如果map 在声明后没有使用复合字面值或者内置的make函数进行 初始化,那么它的值将会是默认值nil
// nil map 可以 遍历(range),可以读取length (len()), 但不可以新增
func main() {
	// 声明一个 map key:string value:int
	var soup map[string]int
	fmt.Println(soup == nil)
	// 对值为nil的map 进行读取 类似 js: var obj={}; obj.onion 不会报错
	measurement, ok := soup["onion"]
	// ok 为false
	if ok {
		fmt.Println(measurement)
	}

	// 并没有执行
	for ingredient, measurement := range soup {
		fmt.Println(ingredient, measurement)
	}

	// 综上 对 nil map 进行 range 读取字段 不会panic
	temperature := map[string]int{
		"Earth": 15,
		"moon":  0,
	}
	temperature["demo"] = 1
	fmt.Println("正常map 可以新增:", temperature)

	// nil map 不可以 新增操作, 会引发 panic
	// soup["demo"] = 1
	// fmt.Println("nil map 新增:", soup)
	fmt.Println("map---", len(soup))
}
