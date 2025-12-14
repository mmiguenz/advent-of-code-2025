package day4_test

import (
	"advent-of-code-2025/day4"
	"testing"
)

func TestCountAccessibleRolls(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rollsGrid []string
		want      int64
	}{
		{
			"sample 1",
			[]string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day4.CountAccessibleRolls(tt.rollsGrid)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("CountAccessibleRolls() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxRollsCanBeRemoved(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		rollsGrid []string
		want      int64
	}{
		{
			"sample 1",
			[]string{
				"..@@.@@@@.",
				"@@@.@.@.@@",
				"@@@@@.@.@@",
				"@.@@@@..@.",
				"@@.@@@@.@@",
				".@@@@@@@.@",
				".@.@.@.@@@",
				"@.@@@.@@@@",
				".@@@@@@@@.",
				"@.@.@@@.@.",
			},
			43,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day4.MaxRollsCanBeRemoved(tt.rollsGrid)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("MaxRollsCanBeRemoved() = %v, want %v", got, tt.want)
			}
		})
	}
}
