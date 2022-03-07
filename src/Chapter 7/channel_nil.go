package main

import (
	"fmt"
	"math/rand"
	"time"
)

// nil 通道

/*
* 1.如果不使用 make初始化通道,那么通道变量的值就是 nil (零值); 意思没有给channel 预分配内存,channel就是 nil channel
* 2. 对nil通道进行发送或接收 不会引起panic,但会导致永久阻塞,阻塞读写
* 3. 对nil通道执行close函数,那么会引起 panic
* 4. nil通道用处:
	 对于包含 select语句的循环,如果不喜欢每次循环都等待select所涉及的多有通道,
	 那么可以先将某些通道设为nil,等到发送值准备就绪之后,再将通道变为一个非nil值并执行发送操作
*/

func main() {
	c := make(chan int)
	go add(c)
	go send(c)
	// 给3秒时间让前俩个 goroutinue 有足够时间运行
	time.Sleep(3 * time.Second)
}

// 不断向 channel c 中发送 [0,2}的随机数
func send(c chan int) {
	for {
		c <- rand.Intn(2)
	}
}

func add(c chan int) {
	sum := 0
	t := time.After(1 * time.Second)

	for {
		select {
		case input := <-c:
			// 不断读取c中的随机数进行家累加
			sum += input
		case <-t:
			fmt.Printf("time.After 返回的类型 是 %T\n", t)
			c = nil
			fmt.Println(sum)
		}
	}
}
