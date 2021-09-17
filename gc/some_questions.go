package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

//尽管 Go 团队宣称 STW 停顿时间得以优化到 100 微秒级别，但这本质上是一种取
//舍。原本的 STW 某种意义上来说其实转移到了可能导致用户代码停顿的几个位置；除
//此之外，由于运行时调度器的实现方式，同样对 GC 存在一定程度的影响。
//目前 Go 中的 GC 仍然存在以下问题：
//1. Mark Assist 停顿时间过长

const (
	windowSize = 200000
	msgCount   = 1000000
)

var (
	best    = time.Second
	bestAt  time.Time
	worst   time.Duration
	worstAt time.Time
	start   = time.Now()
)

func main() {
	f, _ := os.Create("/Users/ahua/goProject/go-notes/gc/questions.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	for i := 0; i < 5; i++ {
		measure()
		worst = 0
		best = time.Second
		runtime.GC()
	}
}

func measure() {
	var c channel
	for i := 0; i < msgCount; i++ {
		c.sendMsg(i)
	}
	fmt.Printf("Best send delay %v at %v,worst send delay: %v at %v. Wall clock: %v\n",
		best, bestAt.Sub(start), worst,
		worstAt.Sub(start), time.Since(start))
}

type channel [windowSize][]byte

func (c *channel) sendMsg(id int) {
	start := time.Now()
	// 模拟发送
	(*c)[id%windowSize] = newMsg(id)
	end := time.Now()
	elapsed := end.Sub(start)
	if elapsed > worst {
		worst = elapsed
		worstAt = end
	}
	if elapsed < best {
		best = elapsed
		bestAt = end
	}
}

func newMsg(n int) []byte {
	m := make([]byte, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}
