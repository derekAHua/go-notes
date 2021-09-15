package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func PrintGC() {
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
