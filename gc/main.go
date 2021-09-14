package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/trace"
	"time"
)

func main() {
	go printGCStatus()

	for i := 0; i < 10; i++ {
		allocate()
		time.Sleep(time.Second * 2)
	}

}

func printGCStatus() {
	t := time.NewTicker(time.Second)
	s := debug.GCStats{}
	m := runtime.MemStats{}
	for {
		select {
		case <-t.C:
			debug.ReadGCStats(&s)
			runtime.ReadMemStats(&m)
			fmt.Printf("gc %d; last@%v; PauseTotal %v; nest_heap_size %v; MSpanInuse %v; HeapObjects %v\n", s.NumGC, s.LastGC, s.PauseTotal, m.NextGC, m.MSpanInuse, m.HeapObjects)
		}
	}
}

// 运行完毕后执行：go tool trace gc/trace.out
func mainByGoToolTrace() {
	f, err := os.Create("gc/trace.out")
	if err != nil {
		panic("open file error !")
	}
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 1000; i++ {
		allocate()
	}
}

func allocate() {
	_ = make([]byte, 2*1<<20)
}
