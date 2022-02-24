package main

import "fmt"

// 结构类型 struct

// Go 允许你将不同类型的东西组合在一起 类似: 类
// 可以声明变量 也可以声明 类型; 类型的好处是 可以复用
// type location struct {
// 	lat, long float64
// }

type location struct {
	lat  float64
	long float64
}

// 别名
type dis location

func distance(loc1, loc2 location) dis {
	// return dis{0.0, 0.0} //初始化赋值
	fmt.Println("loc1--", loc1, loc2)

	return dis{lat: loc1.lat, long: loc2.long}

}
func main() {
	a := location{lat: 3.3}
	b := location{long: 6.6}
	fmt.Println("distance(a, b)", distance(a, b))

	// 符合字面值初始化
	curiosity := location{-4.5895, 137.4417}
	fmt.Println("第一种声明: 按字段定义的顺序进行初始化:", curiosity)

	opportunity := location{lat: -4.5895, long: 137.4417}
	fmt.Println("第二种声明: 成对的字段和值进行初始化:", opportunity)

	// 打印 strcut

	// 第二种 struct 声明 类型: (简写)
	type location2 struct {
		lat, long float64
	}

	cur := location2{-4.5895, 137.4417}
	fmt.Printf("只能打印值:%v\n", cur)
	fmt.Printf("+v可以把struct的花括号字段名都打印出来:%+v\n", cur)

	// struct 复制 (值类型) a 复制 b 改变a; b不会受影响; 传递函数也不会变化

	c := location{lat: 1.123, long: 2.21331}
	d := c          // c 赋值给 d
	d.lat += 3.1415 // 改变 d
	fmt.Println(c, d)

	// 由结构体组成的slice
	type productType struct {
		name string
		age  byte
		lat  float64
		long float64
	}

	// 正常的float64 切片
	lats := []float64{-4.95, 3.74, 9.82}

	// 结构体组合的切片
	locations := []productType{
		{name: "dave", age: 20, lat: 20.123, long: -18.72},
		{name: "Jone", age: 75, lat: 321.02, long: -72.81},
	}
	fmt.Println(lats, locations)

	// 将Struct 编码为 JSON
	// JSON (js object Notation js对象表示法)
	// Go 中 json包的Marshal 函数可以将struct 中的数据 转化为 JSON格式
}
