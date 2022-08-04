package main

import (
	"fmt"
	"sync"
)

//
func main() {
	TestCloseChannel()
}

func TestCloseChannel() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}

// 提供者,生产者

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 5; i++ {
			// fmt.Println("i", i)
			ch <- i
		}
		// 输出完数据 !!!! 就关闭channel
		close(ch)
		wg.Done()
	}()
}

// 接收者;消费者
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			// 这里去接受 当channel被关闭, 即便这里为阻塞等待状态,ok都会为关闭false
			// 返回俩个值, ok为true 代表接受到值, false 代表
			if data, ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("break")
				break
			}
		}
		wg.Done()
	}()
}
