package main

import "testing"

/*
go test -bench='.'
go test -bench='.' -cpuprofile='cpu.prof' -memprofile='mem.prof'

go tool pprof cpu.prof
go tool pprof mem.prof
hide=runtime
top
list TwoSumWithTwoPassHashTable
*/

func BenchmarkTwoSumWithBruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumWithBruteForce([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumWithTwoPassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumWithTwoPassHashTable([]int{2, 7, 11, 15}, 9)
	}
}

func BenchmarkTwoSumOnePassHashTable(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TwoSumOnePassHashTable([]int{2, 7, 11, 15}, 9)
	}
}
