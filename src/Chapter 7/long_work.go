package main

import (
	"fmt"
	"time"
)

// 事件循环
// 中心循环

// Go通过提供的goroutine 作为核心概念,消除了对中心循环的需求,每一个goroutine 就是一个事件循环

// 工作进程(worker)
// 通常会被写成 包含 select语句的for循环:

func worker() {
	next := time.After(time.Second)
	n := 0
	timeout := time.After(3 * time.Second)
	for {
		select {
		case <-next:
			fmt.Println(n)
			n++
			// 	更新
			next = time.After(time.Second)
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}

	}
}

func main() {
	go worker()
	time.Sleep(4 * time.Second)

}
