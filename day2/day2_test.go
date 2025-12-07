package day2_test

import (
	"advent-of-code-2025/day2"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindInvalidIds(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		from int64
		to   int64
		want []int64
	}{
		{
			name: "sample 1",
			from: 11,
			to:   22,
			want: []int64{11, 22},
		},
		{
			name: "sample 2",
			from: 95,
			to:   115,
			want: []int64{99, 111},
		},
		{
			name: "sample 3",
			from: 1188511880,
			to:   1188511890,
			want: []int64{1188511885},
		},
		{
			name: "sample 4",
			from: 1698522,
			to:   1698528,
			want: []int64{},
		},
		{
			name: "sample 5",
			from: 998,
			to:   1012,
			want: []int64{999, 1010},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day2.FindInvalidIds(tt.from, tt.to)

			assert.Equal(t, tt.want, got, "FindInvalidIds()")
		})
	}
}

func TestAddingUpInvalidIds(t *testing.T) {
	tests := []struct {
		name      string
		idsRanges string
		want      int64
	}{
		{
			name:      "sample 1",
			idsRanges: "1188511880-1188511890",
			want:      1188511885,
		},
		{
			name:      "sample 2",
			idsRanges: "1188511880-1188511890,1698522-1698528",
			want:      1188511885,
		},
		{
			name:      "sample 3",
			idsRanges: "998-1012",
			want:      1010 + 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day2.AddingUpInvalidIds(tt.idsRanges)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("AddingUpInvalidIds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsASequenceRepeatedAnyTimes(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		currentNumber string
		want          bool
	}{
		{
			"9 reapeated three times",
			"999",
			true,
		},
		{
			"24 reapeated 5 times",
			"2424242424",
			true,
		},
		{
			"non repeated pattern",
			"1001",
			false,
		},
		{
			"10 repeated twice",
			"1010",
			true,
		},
		{
			"1 repeated 3 times",
			"111",
			true,
		},
		{
			"824 repeated 3 times",
			"824824824",
			true,
		},
		{
			"non repeated 2",
			"8248248245",
			false,
		},
		{
			"non repeated 3",
			"998",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := day2.IsASequenceRepeatedAnyTimes(tt.currentNumber)
			
			if got != tt.want{
				t.Errorf("IsASequenceRepeatedAnyTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
