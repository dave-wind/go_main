package main

import "fmt"

// Go 语言中 没有 class 类 没有 对象 没有 继承
// 但是 struct 和 方法 可以实现 类

// coordinate in degrees, minutes, seconds in a n/s/E/w hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

// 将方法 关联到struct , 且方法可以被关联到你声明的类型上
// 实现一个类如下:
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	// decimal converts a d/m/s coordinate to decimal degrees.
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {

	// Bradbury Landing: 4°35'22.2”s，137°26'30.1"E
	lat := coordinate{4, 35, 22.2, 's'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
