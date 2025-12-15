package day5

import (
	"strconv"
	"strings"
)

func CountFreshIngredients(freshRanges []string, ingredients []int64) int64 {
	freshCount := int64(0)
	for _, ingredient := range ingredients {
		if isFresh(freshRanges, ingredient) {
			freshCount++
		}
	}

	return freshCount
}

func isFresh(freshRanges []string, ingredient int64) bool {
	for _, freshRange := range freshRanges {
		lower, upper := parseRange(freshRange)
		if ingredient >= lower && ingredient <= upper {
			return true
		}
	}
	return false
}

func parseRange(freshRange string) (int64, int64) {
	freshRanges := strings.Split(freshRange, "-")
	lower, _ := strconv.ParseInt(string(freshRanges[0]), 10, 64)
	upper, _ := strconv.ParseInt(string(freshRanges[1]), 10, 64)

	return lower, upper

}
