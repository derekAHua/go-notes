package main

import "testing"

// 查看GC信息
func TestPrintGC(t *testing.T) {
	PrintGC()
}

// 通过go tool分析GC
// 运行完毕后执行：go tool trace gc/trace.out
func TestMainByGoToolTrace(t *testing.T) {
	ByGoToolTrace()
}

// 查看Trace And Heap
//1. 访问(记录20s)：http://127.0.0.1:6060/debug/pprof/trace?seconds=20
//2. go tool trace 生成的文件

//1. curl -sK -v http://localhost:6060/debug/pprof/heap > heap.out
//   go tool pprof heap.out
//1. go tool pprof http://localhost:6060/debug/pprof/heap
//2. 输入web
func TestSeeTraceAndHeap(t *testing.T) {
	SeeTraceAndHeap()
}

func TestSeeTraceAndHeapChange(t *testing.T) {
	SeeTraceAndHeapChange()
}
