package day3

import (
	"strconv"
)

func SumAllMaxJoltage(banks []string) int64 {
	totalSum := int64(0)
	for _, bank := range banks {
		totalSum += MaxJoltage(bank)
	}
	return totalSum
}


func MaxJoltage(bank string) int64 {
	maxNumber := int64(0)
	for i := 0; i < len(bank); i++ {
		for j := i + 1; j < len(bank); j++ {
			number, _ := strconv.ParseInt(string(bank[i])+string(bank[j]), 10, 64)
			
			if number > maxNumber {
				maxNumber = number
			}
		}
	}

	return maxNumber
}
