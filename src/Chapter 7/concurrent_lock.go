package main

import (
	"fmt"
	"strconv"
	"sync"
)

// 共享值 : 类似共享电话
// 竞争条件 (race condition) 不同的goroutine 修改共享值 就会触发竞争条件,报错问题

// 互斥锁 mutex = mutual exclusive
// 方法:
// Lock() Unlock(): 访问共享值的代码必须先锁定互斥锁,然后执行所需操作,操作完成后必须解除互斥锁,不然就会引发竞争条件

// 互斥锁 隐患: 死锁
// 规则:
// 1.尽可能 简化互斥锁保护的代码
// 2. 对每一份共享状态 只使用 一个互斥锁

var mu sync.Mutex

type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

func (v *Visited) VisitLink(url string) int {
	// 一旦被锁住 就会更新映射
	v.mu.Lock()
	defer v.mu.Unlock()
	// 更新
	count := v.visited[url]
	num, err := strconv.Atoi(url)
	if err != nil {
		return 0
	}
	count += num * 2
	// 更新
	v.visited[url] = count
	return count
}
func main() {
	demo1 := Visited{visited: make(map[string]int)}
	demo1.VisitLink("0")
	demo1.VisitLink("3")
	demo1.VisitLink("2")
	demo1.VisitLink("1")

	// time.Sleep(time.Second)

	fmt.Println(demo1.visited)
}
