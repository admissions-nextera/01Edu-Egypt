package main

import "testing"

// Benchmark passing the 8KB struct by value
func BenchmarkProcessValue(b *testing.B) {
	large := LargeStruct{}
	for i := 0; i < b.N; i++ {
		processValue(large)
	}
}

// Benchmark passing the 8KB struct by pointer
func BenchmarkProcessPointer(b *testing.B) {
	large := LargeStruct{}
	for i := 0; i < b.N; i++ {
		processPointer(&large)
	}
}
