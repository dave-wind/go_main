package main

import (
	"fmt"
)

// 结构体
type Item struct {
	ID    int
	Type  string
	Child *Item // 依旧是指针类型 Item
}

// 定义一个接口类型 实际上 结构体 满足了 该接口类型
type ItemClassifier interface {
	IsDoll() bool
}

// 结构体 Item 上定义了一个 IsDoll 方法
func (it *Item) IsDoll() bool {
	fmt.Printf("*it: %v\n", *it)
	if it.Type == "doll" {
		return true
	}
	return false
}

// 定义一个递归函数, 参数: Item类型 返回值也是Item类型
func findDiamond(item Item) Item {
	if item.IsDoll() {
		// 因为Child 是Item指针类型,所以需要解引用,取值传递
		return findDiamond(*item.Child)
	}
	return item

}

func main() {
	doll := Item{
		ID:   1,
		Type: "doll",
		Child: &Item{ // Item 内存地址
			ID:   2,
			Type: "doll",
			Child: &Item{
				ID:   3,
				Type: "doll",
				Child: &Item{
					ID:    4,
					Type:  "diamond",
					Child: nil,
				},
			},
		},
	}
	diamond := findDiamond(doll)
	// findDiamnd()
	fmt.Printf("Item %d is diamond\n", diamond.ID) // 4
}
