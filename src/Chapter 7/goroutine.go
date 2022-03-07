package main

import (
	"fmt"
	"time"
)

// gopher

// 解析:
// 在go语言中 独立的任务叫做 goroutine
//  虽然goroutine 与其他语言中的呢 协程,进程 线程 都有相似之处,但goroutine 和它们并不完全相同
// Goroutine 创建效率非常高
// Go 能直接了当的协同多个并发 (concurrent) 操作

// 启动:
// 只需在调用前面加一个go关键字

func main() {
	for i := 0; i < 5; i++ {
		// 另开辟一个分支线路 执行
		go sleepGopher(i)
		// main 函数结束 就会立即停止所有 goroutine 所以需要等待
	}
	time.Sleep(4 * time.Second)
}

// goroutine 的参数:
//  传递参数和向函数传递参数一样,参数都是按值传递(传入的是副本)
func sleepGopher(id int) {
	// 睡了三秒
	time.Sleep(3 * time.Second)
	fmt.Println("...snore...", id)
}

//注:
/*
1.每次使用go关键字 都会产生一个新的 goroutine
2.表面上看 goroutine似乎在同时运行,但由于计算机处理单元有限,其实技术上来说.这些goroutine不是真的在同时运行
3.计算机处理器 会使用 ”分时“技术 在多个goroutine 上轮流花费一些时间
4.在使用goroutine时,各个goroutine的执行顺序无法确定
*/

// 思考:这段代码有问题,需要白白浪费一秒, 我们需要让代码知道 goroutine 在什么地方结束,所以需要 channel
