package main

/*
。当 goroutine 在等待通道的发送或接收时，我们就说它被阻塞了。
。除了 goroutine 本身占用少量的内存外，被阻塞的 goroutine 并不消耗任何其它资源。
·goroutine 静静的停在那里，等待导致其阻塞的事情来解除阻塞。
。当一个或多个 goroutine 因为某些永远无法发生的事情被阻塞时，
我们称这种情况为死锁。而出现死锁的程序通常会崩溃或挂起。。
*/
func main() {
	c := make(chan int)
	//  all goroutines are asleep - deadlock
	//解决方法: 赋值即可
	go func() { c <- 2 }()
	<-c
}

// 因为创建channle之后,没有给goroutine 给它传值, 通道就会被阻塞
