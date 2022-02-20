package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func randNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func isFixed(val float64) float64 {
	// return math.Trunc(val*1e2+0.5) * 1e-2
	val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return val
}

func Demo() {
	arr := [...]float64{0.05, 0.1, 0.25}

	var result float64 = 0
	times := 0
	for result < 20 {
		num := randNum(0, 2)
		result += arr[num]
		times += 1
	}
	fmt.Println("how many times add ---", times, isFixed(result))

	pig := 0.1
	pig += 0.2
	fmt.Println("pig与0.3 是否相等 要用 math.Abs 结果:", math.Abs(pig-0.3) < 0.0001)

}

func main() {

	// fmt.Println(math.Abs(20.0001-20) == 0)
	Demo()

}
