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

func Benchmark_conPlus(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		conPlus(params)

	}
	b.StopTimer()

}

func Benchmark_conBytes(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		conBytes(params)

	}
	b.StopTimer()

}

func Benchmark_conStrings(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		conCopy(params)

	}
	b.StopTimer()

}
