package main

import (
	"fmt"
	"time"
)

// 计时操作
func main() {
	tick := time.Tick(time.Second)
	i := 1
	for range tick {
		fmt.Println(i)
		i++
	}
}
