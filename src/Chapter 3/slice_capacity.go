package main

import "fmt"

// cap ：容量 如果超出 就按之前的容量*2 扩容
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v,capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	// append 方法
	d1 := []string{"a", "b", "c", "d", "f"} // d1切片长度为5
	d2 := append(d1, "1")                   // 发现长度不够 所以 扩容为 10
	d3 := append(d2, "2", "3", "4")         // 发现容量够 直接添加

	// 综上 d2 d3 对应的底层数组指向的是同一个 要修改 d3[1], d[2] d[3] 都会变！！！
	d3[1] = "Pluto"

	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(d3)

	fmt.Println("-------------------------------------------------------")

	// 三个索引切片操作：
	planets := []string{
		"Mercury", "venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// 该题目， 用第三个索引 进行切片操作 是为了改变 指向的底层数组
	terrestrial := planets[0:4:4] // terrestrial 长度为 4 容量为4 脱离了 planets 指向的是新的底层数组

	// append 往切片内 添加一个元素 返回新切片
	worlds := append(terrestrial, "哈哈") // 原来容量只有4 现在又要添加一个 所以新分配一个数组 容量变为8

	dump("planets", planets)

	dump("terrestrial", terrestrial)

	// worlds
	dump("worlds", worlds)

	//注： 这上面三个底层数据都不一样、

	fmt.Println("------------------------内置make-------------------------------")

	// make 函数对 slice进行预切片
	// 好处：当slice的容器不足还要进行append的时候 Go必须 重新分配新的数组 并对旧值进行copy，
	// 但 通过内置的make函数 可以对slice进行预分配策略，尽量避免额外的内存分配 和  数组复制操作！！

	// 参数：第一个参数：类型；2: len 3:容量 ; 只写第二个 ：代表 长度 容量 都是 第二个参数
	dwarfs := make([]string, 0, 10)
	// 	dwarfs := make([]string, 10)

	dump("dwarfs", dwarfs)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")

	dump("dwarfs", dwarfs)

	// make Demo:
	str := []string{"Dave", "Jones"}
	// demo := terminarl("new", "Dave", "Jones")

	// 解构 展开 str 切片 类似js的 {...obj}
	demo := terminarl("New", str...)
	fmt.Println("terminarl 结果:", demo, "容量：", cap(demo), "长度：", len(demo))
}

func terminarl(prex string, world ...string) []string {
	fmt.Println("world--", world)

	// 为了避免直接修改 world 导致 str 改变，所以用make 生成了一个切片
	var newWords = make([]string, len(world))

	for i := range world {
		newWords[i] = prex + "_" + world[i]
	}
	return newWords
}

// ...的作用：
// 1.做函数的可变参数，2.自动识别数组长度 3.切片的展开 4. ..interface{} 任意类型转换

