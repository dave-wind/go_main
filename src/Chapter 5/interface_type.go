package main

import (
	"fmt"
	"strings"
)

// 为了 复用 通常会把 接口声明为类型
// 按约定: 接口名字 通常以er结尾

// 接口类型优点:
// 1.可以用在任何其他类型(需满足接口) 所使用的地方
// 2.再扩张 接口类型的时候, 依赖接口的方法也会自适应 比如这里的 shout
type talker interface {
	talk() string
	// say() string
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("pew:", int(l))
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

func main() {
	shout(martian{})
	shout(laser(2))

	curiosity := location{-4.5895, 137.4417}
	// 因为在 location 类型上 定义了 String 方法
	// 满足了 fmt Stringer接口 即可当作参数传递给fmt.Println
	fmt.Println(curiosity)

}

// 满足接口
/*
   ·Go 标准库导出了很多只有单个方法的接口。
   Go 通过简单的、通常只有单个方法的接口......来鼓励组合而不是继承，这些接口在各个组件之间形成了简明易懂的界限。
   --- Rob Pike
   标准库常用接口还包括: io.Reader io.Writer json.Marshaler
*/

// 例如 fmt 包声明的 Stringer 接口:
type Stringer interface {
	String() string
}

// location with a latitude, longitude in decimal degrees.
type location struct {
	lat, long float64
}

// location 定义了 String 方法
func (l location) String() string {
	return fmt.Sprintf("%v，%v", l.lat, l.long)
}
