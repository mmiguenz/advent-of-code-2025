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

func TestMaxJoltageBruteForce(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank              string
		amountOfBatteries int
		want              int64
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
			got := day3.MaxJoltageBruteForce(tt.bank, tt.amountOfBatteries)

			if got != tt.want {
				t.Errorf("MaxJoltageBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxJoltageGA(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank              string
		amountOfBatteries int
		want              int64
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
		{
			name:              "sample 5",
			bank:              "565656565656",
			amountOfBatteries: 6,
			want:              666666,
		},
		{
			name:              "sample 6",
			bank:              "121212121212",
			amountOfBatteries: 6,
			want:              222222,
		},
		{
			name:              "sample 7",
			bank:              "9876543210",
			amountOfBatteries: 5,
			want:              98765,
		},
		{
			name:              "sample 8",
			bank:              "111122223333",
			amountOfBatteries: 6,
			want:              223333,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day3.MaxJoltageGA(tt.bank, tt.amountOfBatteries)

			if got != tt.want {
				t.Errorf("MaxJoltageGA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxJoltageStacking(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		bank              string
		amountOfBatteries int
		want              int64
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
			got := day3.MaxJoltageStacking(tt.bank, tt.amountOfBatteries)

			if got != tt.want {
				t.Errorf("MaxJoltageStacking() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumAllMaxJoltage(t *testing.T) {
	tests := []struct {
		name             string
		banks            []string
		amountOfBateries int
		want             int64
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
