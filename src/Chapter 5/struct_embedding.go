package main

import "fmt"

// struct 嵌入 优点:
// 1.实现方法的转发
// 2. 字段 也会被转发
// 3. 结构体可以嵌入任何类型 包括内置类型 string int

type report struct {
	// string
	sol
	temperature // 不写字段名 只写类型 这叫 struct 嵌入
	location    // 好处在于 report 可以直接调用 这些类型 附加的方法
}

type sol int

func (s1 sol) days(s2 sol) int {
	days := int(s2 - s1)
	if days < 0 {
		days = -days // 负负得正 😂
	}
	return days
}

// report 声明的优先级更高
func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
	sol       int
}
type celsius float64 //摄氏度

// 定义一个函数 t为接收者
func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {
	// 只需定义 location
	bradbury := location{-4.5895, 137.4417, 10}
	// 只需定义 temperature
	t := temperature{high: -1.0, low: -78.0}

	report := report{
		sol:         15,
		temperature: t,
		location:    bradbury,
	}

	// report 想调用 average 改造:
	fmt.Println("struct 嵌入 方法转发:", report.average())
	fmt.Println("struct 嵌入 字段也可转发 :", report.temperature.high)
	fmt.Println("struct 嵌入 甚至直接访问 report.high:", report.high)

	fmt.Println("sol 方法转发:", report.days(10))

	fmt.Println("命名冲突--------------------------------")

	// report 不定义 sol字段 就会没有
	// fmt.Println("嵌入 转发优先级", report.sol)

	// 总结:方法可以直接转发 调用 但是字段 需要一层层调用
	fmt.Println("变量自身声明的字段优先级也是最高的,找不到就按类型默认值空值,int 就是0", report.sol)

	// 方法 优先级 :
	fmt.Println("变量自身声明的方法优先级最高:", report.days(2))

	// 继承还是 组合:
	// 优先使用对象组合 而不是类的继承 -----《Go语言设计模式》
	fmt.Println("命名冲突--------------------------------")
}
