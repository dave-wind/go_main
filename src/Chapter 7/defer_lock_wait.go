package main

import (
	"fmt"
	"sync"
)

/*
* wait:
 为了替代 time.Sleep 这样的方法 对于每一个启动的协程, 我们可以先 wg.Add(1),等结束了 再 Done
 初始化执行 用 Wait()
*/
func main() {
	// 锁
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			// 函数最后执行 解锁释放
			defer func() {
				mut.Unlock()
			}()
			// 先锁上
			mut.Lock()
			// 共享内存机制
			counter++
			wg.Done()
		}()
	}
	wg.Wait()
	// time.Sleep(1 * time.Second)
	// 5000; 因为 5000个goroutine 都在共享 counter变量, 所以需要🔒
	fmt.Println(counter)

}
