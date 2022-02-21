package main

import (
	"fmt"
	"unsafe"
)

/*
*@description 整数类型 包含有符号和没符号  取值范围 有8种 int数据类型一共10种 int 和 uint 没有默认是代表多少位不是统称
他们值 取值范围不相同 与架构无关: 可以用 unsafe.Sizeof 取 声明的整数类型 字节数 字节数多少与取值范围无关
int8 -128 to 127
uint8 0 to 255 // uint 表示 没符合只有 正数 和 0
int16 -32768 to 32767
uint16 0 to 65535
int32 -2147483648 to 2147483647  (20亿)
uint32
int64
uint64 // 是整数取值范围最大的数
*/

/*
* int uint 针对目标设备优化的类型
在树莓派2 比较老的移动设备 int 和 uint 都是 32位
在新计算机上 int 和 uint 都是64位
*/
func main() {
	year := 2022
	fmt.Printf("Type %T for %[1]v\n", year) // %T 打印数据类型

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	b := 28
	fmt.Printf("Type %T for %[1]v\n", b)

	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)

	d := true
	fmt.Printf("Type %T for %[1]v\n", d)

	// 颜色 用 unit8来表示 0-255 可以节省很多空间
	var red, green, blue uint8 = 0, 141, 213
	fmt.Println(red, green, blue)
	// 16进制  0,1,2,3,4,5,6,7,8,9,A,B,C,D,E,F 逢16进
	// 16进制 转 10进制 很好转 栗子: 0xAF ==  10*16^1(16的1次方) + 15*16^0 (16的0次方)= 175
	// 10进制 转 16进制  175除以16取余数, 然后商继续 除16 直到商为0结束  175/16 余数15(作为第一位) 商为 10; 10继续除16 商为0 余数为10 (作为第二位)答案:0xAF

	fmt.Println("打印 十六进制--------:")
	// 使用 %x 格式化动词 %02x 02代表不够宽度为2的用0来填充
	fmt.Printf("color: #%02x%02x%02x;\n", red, green, blue)

	fmt.Println("整数 环绕 --------uint8的整数类型 ++ 就会变成0:")
	var demo uint8 = 255 // 0-255
	demo++
	fmt.Println("demo:", demo) // 0 原理如下:
	// 整数环绕原理 栗子: 255++ 的问题 (uint8) uint8 代表最大八位
	/*
		*   11111111
		*  +00000001
		    --------
		   100000000
		   本来++ 是进一位的 但是uint8声明的变量 内存只能是八位 所以就变0咯 😂
	*/

	var t int16 = 127
	fmt.Println("demo内存占多少字节---", unsafe.Sizeof(demo))
	fmt.Println("t int16 127 内存占多少字节---", t, unsafe.Sizeof(t))

	fmt.Println("打印每个bit:-----------") // 使用%b格式化

	var bb uint8 = 3
	fmt.Printf("bb=3 的bit: %08b\n", bb)
	bb++
	fmt.Printf("bb++后: %08b\n", bb)

}
