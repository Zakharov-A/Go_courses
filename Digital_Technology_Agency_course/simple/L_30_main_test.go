package main

import (
	"testing"
)

func Benchmark_conCopy(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		conCopy(params)

	}
	b.StopTimer()

}
