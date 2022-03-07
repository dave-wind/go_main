package main

import (
	"fmt"
	"time"
)

// 通道(channel)
// 1.可以在多个goroutine 之间安全的传值
// 2.可以用作变量 函数参数 结构体字段...
// 3. 创建: 用make 函数,并指定其传输数据类型即可:
/*
  c:= make(chan int)
*/

// 4. 通道channel 发送 ,接受:
// 使用左箭头操作符 <- 向通道发送值 或 从通道接受值:
/*
* 向通道发送值: c<-99
* 向通道接收值: r:= <-c
 */
func main() {
	c := make(chan int)
	// 一共 创建了 5个 goroutine, 用 创建的channel,把 main 和 5个 goroutine 连接起来
	for i := 0; i < 5; i++ {
		go sleepGopher(i, c)
	}
	for i := 0; i < 1; i++ {
		// 先一直等待,等到了接受通道里的值,就执行
		// 一共执行五次
		gopherID := <-c
		fmt.Println("gopher:   ", gopherID, "has finished sleeping")
	}

	// 内置函数 close 关闭 channel
	close(c)
}

func sleepGopher(id int, c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
	// id 值 传到通道里
	c <- id
}

// 局限性:

// 以上这样的情况,只能多个goroutine 之间传递是同一种类型的值 时候是好用的,多种不同类型就不可用了
//所以需要 select处理多个通道

// 文章:
// https://segmentfault.com/a/1190000018719303
