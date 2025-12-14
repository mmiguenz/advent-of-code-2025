package day3_test

import (
	"advent-of-code-2025/day3"
	"testing"
)

func TestMaxJoltage(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank string
		want int64
	}{
		{
			"sample 1",
			"987654321111111",
			98,
		},
		{
			"sample 2",
			"811111111111119",
			89,
		},
		{
			"sample 3",
			"234234234234278",
			78,
		},
		{
			"sample 4",
			"818181911112111",
			92,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.MaxJoltage(tt.bank)

			if got != tt.want {
				t.Errorf("MaxJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxJoltageV2(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank string
		amountOfBatteries int
		want int64
	}{
		{
			"sample 1",
			"987654321111111",
			12,
			987654321111,
		},
		{
			"sample 2",
			"811111111111119",
			12,
			811111111119,
		},
		{
			"sample 3",
			"234234234234278",
			12,
			434234234278,
		},
		{
			"sample 4",
			"818181911112111",
			12,
			888911112111,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.MaxJoltageV2(tt.bank, tt.amountOfBatteries)

			if got != tt.want {
				t.Errorf("MaxJoltageV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAllMaxJoltage(t *testing.T) {
	tests := []struct {
		name string 		
		banks []string
		amountOfBateries int
		want  int64
	}{
		{
			"The total output joltage is the sum of the maximum joltage from each bank, so in this example, the total output joltage is 98 + 89 + 78 + 92 = 357",
			[]string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			2,
			357,
		},
		{
			"The total output joltage is now much larger: 987654321111 + 811111111119 + 434234234278 + 888911112111 = 3121910778619",
			[]string{
				"987654321111111",
				"811111111111119",
				"234234234234278",
				"818181911112111",
			},
			12,
			3121910778619,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.SumAllMaxJoltage(tt.banks, tt.amountOfBateries)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("SumAllMaxJoltage() = %v, want %v", got, tt.want)
			}
		})
	}
}
