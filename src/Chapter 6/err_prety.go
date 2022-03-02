package main

import (
	"fmt"
	"io"
	"os"
)

/*
@description 优雅的处理错误
*/
// 定义一个结构体
type safeWriter struct {
	w   io.Writer
	err error // 用err 存储错误
}

// 定义一个方法 接受者是 结构体
func (sw *safeWriter) writeln(str string) {
	// 调用的过程中 有错误 就直接报错
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintf(sw.w, str)
}

func setProverbs(name string) error { // error 内置类型
	// 创建文件
	f, err := os.Create(name)
	// 创建文件出错 报错err
	if err != nil {
		return err
	}
	// 函数结束之前 关闭文件
	defer f.Close()

	// 优雅的 写入
	sw := safeWriter{w: f}
	sw.writeln("Errors are values")
	sw.writeln("Don't communicate by sharing memory, share memory by communicating.\n")
	sw.writeln("The bigger the interface, the weaker the abstraction. interfacef says nothing.\n")

	return sw.err
}

func main() {
	err := setProverbs("prety_err.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
