package day2

import (
	"strconv"
	"strings"
)

// idRange contains the start and end of a range of IDs.
type idRange struct {
	from int64
	to   int64
}

func AddingUpInvalidIds(idsRanges string) int64 {
	ranges := buildRanges(idsRanges)
	invalidIds := []int64{}
	totalSumInvalids := int64(0)

	for _, r := range ranges {
		invalids := FindInvalidIds(r.from, r.to)
		invalidIds = append(invalidIds, invalids...)
	}

	for _, id := range invalidIds {
		totalSumInvalids += id
	}

	return totalSumInvalids
}

func buildRanges(idRangesStr string) []idRange {
	rangePairs := strings.Split(idRangesStr, ",")
	ranges := make([]idRange, 0, len(rangePairs))

	for _, pair := range rangePairs {

		bounds := strings.Split(pair, "-")
		from, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			panic(err)
		}
		to, err := strconv.ParseInt(bounds[1], 10, 64)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, idRange{from: from, to: to})
	}
	return ranges
}

func FindInvalidIds(from int64, to int64) []int64 {
	foundElems := []int64{}
	for i := from; i <= to; i++ {
		currentNumber := strconv.FormatInt(i, 10)
		numberLen := len(currentNumber)
		if numberLen%2 == 0 {
			if currentNumber[:numberLen/2] == currentNumber[numberLen/2:] {
				foundElems = append(foundElems, i)
			}
		}
	}
	return foundElems
}
