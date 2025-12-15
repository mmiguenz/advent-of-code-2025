package day5_test

import (
	"advent-of-code-2025/day5"
	"testing"
)

func TestCountFreshIngredients(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		freshRanges []string
		ingredients []int64
		want        int64
	}{
		{
			"sample 1, only 3 fresh ingredients",
			[]string{
				"3-5",
				"10-14",
				"16-20",
				"12-18",
			},
			[]int64{1, 5, 8, 11, 17, 32},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day5.CountFreshIngredients(tt.freshRanges, tt.ingredients)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("CountFreshIngredients() = %v, want %v", got, tt.want)
			}
		})
	}
}
