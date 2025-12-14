package day3

import (
	"strconv"
	"strings"
)

func SumAllMaxJoltage(banks []string, amountOfBateries int) int64 {
	totalSum := int64(0)
	for _, bank := range banks {
		totalSum += MaxJoltageV2(bank, amountOfBateries)
	}
	return totalSum
}

func MaxJoltage(bank string) int64 {
	maxNumber := int64(0)
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			number := buildNumber(string(bank[i]), string(bank[j]))

			if number > maxNumber {
				maxNumber = number
			}
		}
	}

	return maxNumber
}

func MaxJoltageV2(bank string, amountOfBatteries int) int64 {
	maxNumber := int64(0)

	allPossibleJoltages := combinations(strings.Split(bank, ""), amountOfBatteries)

	for _, combination := range allPossibleJoltages {
		number := buildNumber(combination...)

		if number > maxNumber {
			maxNumber = number
		}
	}

	return maxNumber
}

func buildNumber(digit ...string) int64 {
	digits := strings.Join(append([]string{}, digit...), "")
	number, _ := strconv.ParseInt(digits, 10, 64)
	return number
}

func combinations(elements []string, k int) [][]string {
	n := len(elements)
	if k < 0 || k > n {
		return [][]string{}
	}

	// indexes will store the current combination of positions
	idx := make([]int, k)
	for i := 0; i < k; i++ {
		idx[i] = i
	}

	var result [][]string

	for {
		// Build combination from indexes
		combo := make([]string, k)
		for i, pos := range idx {
			combo[i] = elements[pos]
		}
		result = append(result, combo)

		// Find the rightmost index to increment
		i := k - 1
		for i >= 0 && idx[i] == i+n-k {
			i--
		}

		// If no such index exists, we are done
		if i < 0 {
			break
		}

		// Increment this index
		idx[i]++

		// Reset all following indexes
		for j := i + 1; j < k; j++ {
			idx[j] = idx[j-1] + 1
		}
	}

	return result
}
