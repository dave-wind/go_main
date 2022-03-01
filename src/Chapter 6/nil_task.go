package main

import "fmt"

/*
亚瑟被一位骑士挡住了去路。正如leftHand *item 变量的值为nil所示，这位英雄手上正空无一物。
请实现一个拥有 pickup(i *item)和 give(to *character)等方法的 character 结构，然后使用你在本节学到的知识编写一个脚本，使得亚瑟可以拿起一件物品并将其交给骑士，与此同时为每个动作打印出适当的描述。
*/

type item struct {
	name string
}
type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if c == nil || i == nil {
		return
	}
	fmt.Printf("%v picks up a %v\n", c.name, i.name)
	c.leftHand = i
}
func (c *character) give(to *character) {
	if c == nil || to == nil {
		return
	}
	if c.leftHand == nil {
		fmt.Printf("%v has nothing togive\n", c.name)
		return
	}
	if to.leftHand != nil {
		fmt.Printf("%v's hands are full\n", to.name)
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v give %v a %v\n", c.name, to.name, to.leftHand.name)
}

func (c character) String() string {
	if c.leftHand == nil {
		return fmt.Sprintf("%v is carrying nothing", c.name)
	}
	return fmt.Sprintf("%v is carrying a %v", c.name, c.leftHand.name)
}

func main() {
	arthur := &character{name: "Dave"}

	shrubbery := &item{name: "shrubbery"}
	arthur.pickup(shrubbery)

	knight := &character{name: "Knight"}
	arthur.give(knight)
	fmt.Println(arthur)
	fmt.Println(knight)
}
