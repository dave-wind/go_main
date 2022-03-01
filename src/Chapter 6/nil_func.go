package main

import (
	"fmt"
	"sort"
	"strings"
)

// 当参数为 函数,提供默认行为:
func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		// 返回一个默认行为 (这里为排序)
		less = func(i, j int) bool { return s[i] < s[j] }

	}
	// sor 排序方法
	sort.Slice(s, less)
}

// 方法一: 可以按照字符串从短到长的顺序排序
func sortSingleString(s string) string {
	bt := []byte(s)
	sort.Slice(bt, func(i, j int) bool {
		return bt[i] < bt[j]
	})
	return string(bt)
}

// 方法二: 单个字符串 字母排序
func sortSingleString2(str string) string {
	// 字符串分割,得到字符串切片  js: arr= str.split()
	split := strings.Split(str, "")
	// 对切片进行排序
	sort.Strings(split)
	// 切片 联结成新的字符串
	return strings.Join(split, "")
}

func main() {

	// 声明一个函数类型 的变量 值为nil
	var fn func(a, b int) int
	fmt.Printf("fn--%v\n", fn == nil)

	// 检查函数的数值是否为nil,并在有需要的时候提供默认行为

	food := []string{"b", "a", "c"}

	sortStrings(food, nil)
	fmt.Println(food)

	fmt.Println("字符串排序1:", sortSingleString("food"))
	fmt.Println("字符串排序2:", sortSingleString2("food"))

}
