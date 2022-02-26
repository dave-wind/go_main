package main

import "fmt"

// 组合

/*
在面向对象的世界中，对象由更小的对象组合而成。
术语:对象组合或组合
·Go 通过结构体实现组合(composition)。
组合 可以实现其他语言那种继承
·Go 提供了“嵌入”(embedding)特性，它可以实现方法的转发(forwarding)
*/

// type report struct {
// 	sol int // 火星问题
// 	high,low float64
// 	lat,long float64
// }
// 把这样的类型 拆分如下改写(为了更方便 复用性更高):

type report struct {
	sol         int
	temperature temperature
	location    location
}
type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}
type celsius float64 //摄氏度

// 定义一个函数 t为接收者
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

// 方法的转发 把average 转发到 temperature的average
func (r report) average() celsius {
	return r.temperature.average()
}

func main() {
	// 只需定义 location
	bradbury := location{-4.5895, 137.4417}
	// 只需定义 temperature
	t := temperature{high: -1.0, low: -78.0}

	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° c\n", report.temperature.high)

	fmt.Println("方法组合-------------------------------")

	// temperature 脱离与  report
	fmt.Println("t变量 调用 temperature average:", t.average())

	// 如果 report 想调用 average?
	fmt.Println("report变量 想调用average:", report.temperature.average())

	// report 想调用 average 改造:
	fmt.Println("在变量r 上 定义 方法", report.average())

}
