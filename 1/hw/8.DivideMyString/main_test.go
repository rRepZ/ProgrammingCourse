package main

import "testing"

func BenchmarkConcatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatString()
	}
}

func BenchmarkConcatStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringBuilder()
	}
}

func BenchmarkConcatStringBuilderAllocateMemory(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatStringBuilderAllocateMemory()
	}
}
