package main

import (
	"fmt"
	"os"
)

/*
. Errors are values. solenovex bilibili
Don't just check errors, handle them gracefully Don't panic.
Make the zero value useful.
The bigger the interface, the weaker the abstraction. interfacef says nothing.
Gofmt's style is no one's favorite, yet gofmt is everyone's favorite. Documentation is for users.
A little copying is better than a little dependency. Clear is better than clever. Concurrency is not parallelism.
Don't communicate by sharing memory, share memory by communicating.
Channels orchestrate, mutexes serialize.
*/

/*
。写入文件的时候可能出错:
路径不正确权限不够磁盘空间不足
·文件写入完毕后，必须被关闭，确保文件被刷到磁盘上，避免资源的泄露。
*/

func proverbs(name string) error { // error 内置类型
	// 创建文件
	f, err := os.Create(name)
	if err != nil {
		return err
	}

	/*
		使用 defer 关键字，Go 可以确保所有 deferred 的动作可以在函数返回前执行。(例子 28.4)
		可以 defer 任意的函数和方法。 defer 并不是专门做错误处理的。
		defer 可以消除必须时刻惦记执行资源释放的负担
	*/
	// 无论文件写入是否成功 都要把文件给关闭，否则会引起资源泄露
	defer f.Close()

	// 对文件进行写入
	_, err = fmt.Fprintln(f, "Errors are values.")
	if err != nil {

		return err
	}
	_, err = fmt.Fprintln(f, "Don't just check errors, handle them gracefully.")
	return err
}

func main() {
	err := proverbs("proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
