package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// 题目定义 一个 结构体 输出 json

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%v,%v,%.2f,%v", c.h, c.d, c.m, c.s)
}

// 	将结构体 度分秒 格式坐标 转变为 十进制
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	// decimal converts a d/m/s coordinate to decimal degrees.
	return sign * (c.d + c.m/60 + c.s/3600)
}

type toJSon struct {
	DD  float64 `json:"decimal"`
	DMS string  `json:"dms"`
	D   float64 `json:"degrees"`
	M   float64 `json:"minutes"`
	S   float64 `json:"seconds"`
	H   string  `json:"hemisphere"`
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(toJSon{
		DD:  c.decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

type location struct {
	Name string     `json:"name"`
	Lat  coordinate `json:"latitude"`
	Long coordinate `json:"longitude"`
}

func main() {
	c := coordinate{
		d: 300.00,
		m: 135,
		s: 54.0,
		h: 4,
	}
	fmt.Println(c.String())

	fmt.Println("输出json-------------------")

	elysium := location{
		Name: "Elysium Planitia",
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	bytes, err := json.MarshalIndent(elysium, "", "")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
