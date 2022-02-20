package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()            //获取当前时间
	timestamp1 := now.Unix()     //时间戳(秒 不是毫秒)
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)

	// Unix 系统 时间以1970年1月1日至今的秒数来表示的
	// 避免时间环绕 到2038年 这个数会超过20多亿 超过了int32的范围 所以声明时间 可以用 int64 || uint64
	future := time.Unix(time.Now().Unix(), 0)
	fmt.Println("格式化时间", future)
}
