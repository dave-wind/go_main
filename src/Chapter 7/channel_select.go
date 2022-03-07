package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 使用select 处理多个通道,等待不同类型的值

// Select语句:
// 1.该语句包含的每个case 都持有一个通道,用来发送和接受数据
// 2.select 会等待直到某个case分支操作就绪,然后就会执行该case分支
// 3.注意: 在不包含任何case的情况下将永远等下去!!!
func main() {
	c := make(chan int)
	// 一共 创建了 5个 goroutine, 用 创建的channel,把 main 和 5个 goroutine 连接起来
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	// time.After函数,返回一个通道,该通道在指定时间后会接收到一个值(发送该值的goroutine 是Go运行时的一部分)
	// 类似 web接口里超时处理
	timeout := time.After(2 * time.Second)

	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher:   ", gopherID, "has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			// 这里通过 return 结束函数,从而停止所以goroutine 不然会占内存
			return
		}
		fmt.Println("i", i)
	}
}

func sleepGopher(id int, c chan int) {
	// 1-4 毫秒之间 随机一个数值
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}

// 注:
/*
即使已经停止等待goroutine ,但只要main函数还没有返回,仍在运行的goroutine 将会继续占用内存

select {} // 没有可用 channel，会阻塞 main goroutine。
*/
