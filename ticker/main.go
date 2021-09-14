package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(time.Second)
	i := 1
	for _ = range tick {
		fmt.Println(i)
		i++
	}
}
