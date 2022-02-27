package main

import (
	"fmt"
	"strings"
)

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

//!!!如果 方法接受者 是指针类型 :
type laser int

func (l *laser) talk() string {
	return strings.Repeat("pew", int(*l))
}

func main() {
	shout(martian{})
	// 传入指针 满足接口 也可以
	shout(&martian{})

	pew := laser(2)
	// 指向laser的指针 满足了接口; 方法接受者 是指针类型 传入的需要满足指针
	shout(&pew)
}
