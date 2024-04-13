package main

// $ go test main_test.go -v

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func TestIntMinBasic(t *testing.T) {
// 	ans := IntMin(2, -2)
// 	if ans != -2 {
// 		// t.Error* will report test failures but continue executing the test
// 		// t.Fatal* will report test failures and stop the test immediately
// 		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
// 	}
// }

// Writing tests can be repetitive, so it's idiomatic to use a table-driven style
// Test inputs and expected outputs are listed in a table and a single loop walks over them and performs the test logic

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{0, 0, 0},
		{-1, 0, -1},
		{0, -1, -1},
		{-1, -1, -1},
		{1, 1, 1},
		{-1, 1, -1},
		{1, -1, -1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// $ go test main_test.go -bench=.

// BenchmarkIntMin provides performance data for the IntMin function
// The benchmark function must run the target code b.N times
// During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
