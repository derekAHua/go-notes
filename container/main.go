package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"strconv"
)

// go container
func main() {
	//h() // use heap
	//l() // use list
	r() // use ring 环形链表
}

func r() {
	ri := ring.New(5)
	for i := 0; i < 5; i++ {
		ri.Value = i
		ri = ri.Next()
	}

	ri.Do(func(i interface{}) {
		fmt.Print(i)
	})
	fmt.Println()

	r2 := ri.Move(3) // 移动3个元素
	fmt.Println("元素", r2.Value)
	fmt.Println("------")
	r2.Do(func(i interface{}) {
		fmt.Print(i)
	})
	fmt.Println()

	link := ri.Link(r2) // 移除r-r2个元素
	link.Do(func(i interface{}) {
		fmt.Print(i)
	})
	fmt.Println()

	ri.Do(func(i interface{}) {
		fmt.Print(i)
	})
}

func h() {
}

func l() {
	lis := list.New()
	for i := 0; i < 10; i++ {
		lis.PushBack("a" + strconv.Itoa(i))
	}
	for i := lis.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
