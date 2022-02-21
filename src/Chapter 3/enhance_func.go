package main

import (
	"fmt"
	"math/rand"
)

// 函数声明三种:
// 1.将函数赋给变量
// 2.声明函数类型
// 3.自执行函数
// 4.匿名函数 (就是没有名字的函数,在Go里也称作函数字面值)
// 因为: 函数字面值需要保留外部作用域变量的引用 索引函数字面值都是闭包
// 闭包 (closure): 就是由于匿名函数封闭 并包围作用域中的变量而得名

// 自定义类型
type kelvin float64

// 2栗子: 声明类型为 一个函数类型
type sensor func() kelvin

// 函数声明 = 匿名函数 赋值给一个变量
var realSensor = func() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// 闭包函数 + 函数作为参数 + 函数作为返回体

func calibrate(s sensor, offset kelvin, not bool) sensor {
	return func() kelvin { // 疑问返回的匿名函数 不可以传参数吗?
		// 按值传递
		temp := offset
		if not {
			temp = 0
		}
		return s() + temp
	}
}

// 3栗子:匿名函数
var f = func() {
	fmt.Println("Dress up for masquerade")
}

func main() {
	// 调用匿名函数
	f()

	// 4栗子: 自执行函数
	func() {
		fmt.Println("自执行函数例如 js里的 iiFE")
	}()

	var num kelvin = 10
	// 1栗子: 将函数赋给变量
	sensor := calibrate(realSensor, num, false)
	fmt.Println("sensor() 结果:", sensor())
	fmt.Println("num 结果:", num)

	fmt.Println("------------------------------")

	// 多次调用 sensor 看看是否每次都产生不同的随机数
	sensor2 := calibrate(fakeSensor, 0, true)()
	sensor3 := calibrate(fakeSensor, 0, true)
	sensor4 := calibrate(fakeSensor, 0, true)
	sensor5 := calibrate(fakeSensor, 0, true)
	fmt.Println("随机数结果:", sensor2, sensor3(), sensor4(), sensor5())
}
