package main

import "fmt"

// map 是 Go提供的另外一种集合
// 它 可以将key 映射到value
// 可快速通过 key 找到对应的 value
// 它的 key 几乎可以是任何类型
func main() {

	// 声明map
	temperature := map[string]int{
		"Earth": 15,
		"moon":  0,
	}
	temp := temperature["Earth"]
	fmt.Println(temp)

	// 修改
	temperature["Earth"] = 10
	fmt.Println("修改过的", temperature["Earth"])

	// 获取map上没有的元素 : 默认是int 的 空值:0
	fmt.Println("no key", temperature["moon"])

	// 逗号 与 ok 的用法
	if moon, ok := temperature["moon"]; ok {
		fmt.Printf("ok 是bool 值 为 %v 代表对象有这个key\n", ok)
	} else {
		fmt.Printf("ok 是bool 值 为 %v 代表没有 moon默认:%v \n", ok, moon)
	}

	fmt.Println("----------make-----------------------")

	// make 对 对象 进行预分配
	// make 函数接受 一个或俩个参数,
	// 第二个参数 用于为指定数量的key 预先分配空间
	// make 函数 创建的map 初始化 长度为  0
	obj := make(map[float64]int, 8)

	fmt.Println("make 预分配", obj, "长度:", len(obj), "容量:")

	// map 多值类型
	obj2 := map[string]interface{}{
		"id":   2,
		"name": "demo",
	}
	fmt.Println("map多值类型:", obj2, "长度:", len(obj2))

	// map在以下情况不会被复制:
	// 数组,int,float64 等类型在赋值给新变量或传递至函数/方法的时候 会创建相应的副本
	// 但是 map 就不会!!!! map 赋值 和 作为函数传递过去 是引用 所以改了都会被改

	var demo byte = 1
	fmt.Println("值传递结果", handleDemo(demo))
	fmt.Println("dmeo 值不会变 因为 调用handleDemo 值被copy了", demo)

	// delete 方法
	delete(obj2, "name")
	fmt.Println("obj2 被delete 后", obj2)
}

func handleDemo(num byte) byte {
	num += 1
	return num
}
