package main

import (
	"fmt"
)

// 方法: 属于某个类型 || 与某个类型所关联的 func
func main() {
	// 关键字 type 可以用来声明新类型
	type cwl float64 // 但是其底层类型还是 float64 一般声明的变量不赋值 默认就是float64哦

	var temperature cwl = 20

	// 当然也可以 进行 加减等 也可以像float64 那些使用
	temperature += 10
	fmt.Println("temperature", temperature)

	//  为什么要声明新类型??：极大的提高代码可读性和可靠性,但是不同类型是无法混用的:
	// var warmUp float64 = 10
	// temperature += warmUp // 不同类型不能混用

	var k kelvin = 204
	fmt.Println("k.celsius(1,2)", k.celsius(1, 2))

	// 	声明函数类型
	fmt.Println("声明函数类型:", mesureTemperature(10, demo))
}

/*
*@important!! 通过方法添加行为
在Go里 它提供了 方法 但是没提供类和对象
*/

type kelvin float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius { // kelvinToCelsius 函数
	return celsius(k - 273.15)
}

// 可以将方法与同包中声明的任何类型相关联 但不可以是int float64 等预声明类型进行关联

// 声明 一个 叫 kelvin的类型 关联了一个叫celsius的方法
// 注意⚠️:方法只能有一个接收者 k 也可以作为参数
func (k kelvin) celsius(e kelvin, d kelvin) celsius { // kelvin 类型的 celsius 方法
	return celsius(k - e - d)
}

// 声明函数类型
type sensor func() int

func demo() int {
	return 10
}

// 为函数声明类型 有助于 精简和明确 代码
func mesureTemperature(samples int, s sensor) int {
	return samples + s()
}
