package main

import "fmt"

// coordinate in degrees, minutes, seconds in a n/s/E/w hemisphere.
type coordinate struct {
	d, m, s float64
	h       rune
}

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

type location struct {
	lat, long float64
}

// Go语言 没有专用的构造函数 但 构造函数 规则: new+大写类型名的函数,
// 通常是用来构造数据的,例如:newPerson NewPerson  New 可倒出 (大写)
func newLocation(lat, long coordinate) location {
	return location{lat: lat.decimal(), long: long.decimal()}
}

func main() {

	lat := coordinate{4, 35, 22.2, 's'}
	long := coordinate{137, 26, 30.12, 'E'}

	curiosity := newLocation(lat, long)
	fmt.Println("curiosity:", curiosity)
	fmt.Printf("由构造函数声明的curiosity:%+v \n", curiosity)
}
