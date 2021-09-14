package test

import (
	"fmt"
	"testing"
)

// 批量测试
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%d", i)
		}
	}
}

// go test -v -bench=. bench_test.go

// +时间
// go test -v -bench=. -benchtime=1s bench_test.go

// +内存
// go test -v -bench=BenchmarkAdd -benchmem bench_test.go
