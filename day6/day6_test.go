package day6_test

import (
	"advent-of-code-2025/day6"
	"testing"
)

func TestGrandTotal(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		homeWorkLines []string
		want          int64
	}{
		{
			"33210 + 490 + 4243455 + 401 = 4277556",
			[]string{"123 328  51 64 ","  45 64  387 23 ", "  6 98  215 314", "*   +   *   +  "},
			4277556,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day6.GrandTotal(tt.homeWorkLines)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GrandTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGrandTotalV2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		homeWorkLines []string
		want          int64
	}{
		{
			"1058 + 3253600 + 625 + 8544",
			[]string{"123 328  51 64 "," 45 64  387 23 ", "  6 98  215 314", "*   +   *   +  "},
			3263827,
		},		
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day6.GrandTotalV2(tt.homeWorkLines)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("GrandTotalV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
