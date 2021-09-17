package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

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

// 模拟 Too much GC
// 执行：go test -bench=BenchmarkGCLargeGs -run=^$ -count=5 -v . | tee 4.txt
func BenchmarkGCLargeGs(b *testing.B) {
	wg := sync.WaitGroup{}
	for ng := 100; ng <= 1000000; ng *= 10 {
		b.Run(fmt.Sprintf("#g-%d", ng), func(b *testing.B) {
			// 创建大量 goroutine，由于每次创建的 goroutine 会休眠
			// 从而运行时不会复用正在休眠的 goroutine，进而不断创建新的 g
			wg.Add(ng)
			for i := 0; i < ng; i++ {
				go func() {
					time.Sleep(100 * time.Millisecond)
					wg.Done()
				}()
			}
			wg.Wait()
			// 现运行一次 GC 来提供一致的内存环境
			runtime.GC()
			// 记录运行 b.N 次 GC 需要的时间
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				runtime.GC()
			}
		})
	}
}
