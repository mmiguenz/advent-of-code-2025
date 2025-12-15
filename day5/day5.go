package day5

import (
	"fmt"
	"slices"
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

func CountFreshIngredientsAvailable(freshRanges []string) int64 {
	slices.SortFunc(freshRanges, sortRangesAscendingCriteria)

	unionRanges := unionOverlappedRanges(freshRanges)

	return countFreshItemsInRanges(unionRanges)

}

func countFreshItemsInRanges(freshRanges []string) int64 {
	totalFresh := int64(0)
	for _, rang := range freshRanges {
		lower, upper := parseRange(rang)

		totalFresh += upper - lower + 1
	}

	return totalFresh
}

func sortRangesAscendingCriteria(i, j string) int {
	lower1, _ := parseRange(i)
	lower2, _ := parseRange(j)

	if lower1 < lower2 {
		return -1
	}
	if lower1 > lower2 {
		return 1
	}
	return 0
}

func unionOverlappedRanges(freshRanges []string) []string {
	if len(freshRanges) < 2 {
		return freshRanges
	}
	range1 := freshRanges[0]
	range2 := freshRanges[1]

	range1Lower, range1Upper := parseRange(range1)
	range2Lower, range2Upper := parseRange(range2)

	if range1Upper >= range2Lower {
		newUpper := max(range1Upper, range2Upper)

		return unionOverlappedRanges(append([]string{fmt.Sprintf("%d-%d", range1Lower, newUpper)}, freshRanges[2:]...))
	}

	return append([]string{range1}, unionOverlappedRanges(freshRanges[1:])...)
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
	lower, err := strconv.ParseInt(string(freshRanges[0]), 10, 64)

	if err != nil {
		panic(err)
	}
	upper, err := strconv.ParseInt(string(freshRanges[1]), 10, 64)

	if err != nil {
		panic(err)
	}

	return lower, upper
}
