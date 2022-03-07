package main

import (
	"fmt"
	"strings"
)

//demo:        upstream   downstream           upstream     downstream
// sourceGopher ------> () ------> filterGopher ------> ()  -------->  printGopher

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0, c1)
	printGopher(c1)
}

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apply", "goobye all"} {
		fmt.Println("source----:", v)
		downstream <- v
	}
	close(downstream)

}

func filterGopher(upstream, downstream chan string) {
	// range 作用:从通道里读取值 直到其关闭为止
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			fmt.Println("filter----:", item)
			downstream <- item
		}
	}

	close(downstream)
	// fmt.Println("downstream", downstream)
}

// Go 语言 运行在没有值的情况下通过close函数 关闭通道 close(c)
// 通道被关闭以后无法写入任何值,如果尝试写入将引发panic
// 尝试读取被关闭的通道会获得与通道类型对应的零值
// 执行如下代码可知 通道是否会被关闭:
// v,ok:=<-c; ok为false 证明通道已经关闭了

// func filterGopher(upstream, downstream chan string) {
// 	for {
// 		item, ok := <-upstream
// 		if !ok {
// 			close(downstream)
// 			fmt.Println("关闭", item)
// 			return
// 		}
// 		if !strings.Contains(item, "bad") {
// 			downstream <- item
// 		}
// 	}
// }

func printGopher(upstream chan string) {
	// for {
	// 	v := <-upstream
	// 	if v == "" {
	// 		return
	// 	}
	// 	fmt.Println(v)
	// }
	for v := range upstream {
		fmt.Println(v)
	}
	// time.Sleep(3 * time.Second)
	// close(upstream)
}
