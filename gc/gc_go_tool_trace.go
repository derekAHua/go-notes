package main

import (
	"os"
	"runtime/trace"
)

func ByGoToolTrace() {
	f, err := os.Create("/Users/ahua/goProject/go-notes/gc/trace.out")
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
