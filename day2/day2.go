package day2

import (
	"fmt"
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

		if // The `IsASequenceRepeatedAnyTimes` function is checking if a given string contains a repeated
		// sequence of characters. It iterates through the input string and tries to find patterns that
		// repeat multiple times. It keeps track of visited patterns and their occurrences to determine if
		// a sequence is repeated more than once in the string. The function returns true if it finds a
		// repeated sequence, otherwise false.
		IsASequenceRepeatedAnyTimes(currentNumber) {
			foundElems = append(foundElems, i)
		}

	}
	fmt.Println("Found elements: ", foundElems)
	return foundElems
}

func isASequenceRepeatedTwice(currentNumber string) bool {
	numberLen := len(currentNumber)
	return numberLen%2 == 0 && currentNumber[:numberLen/2] == currentNumber[numberLen/2:]
}

func IsASequenceRepeatedAnyTimes(currentNumber string) bool {
	pattern := ""
	visited := map[string]int64{}
	lenPatternFirstVisited := map[int]int64{}
	p0 := 0
	p1 := 0

	for p1 < len(currentNumber) {
		pattern = currentNumber[p0 : p1+1]

		lenCurrentPattern := len(pattern)
		visits := visited[pattern]

		if visits == 0 {
			if lenPatternFirstVisited[lenCurrentPattern] == 0 {
				lenPatternFirstVisited[lenCurrentPattern] = 1
				p0, p1 = movePositionToNextPattern(p0, p1, currentNumber, lenCurrentPattern)

			} else {
				p0, p1 = restartPatternPosition(p0)
			}

			visited[pattern] = 1

		} else {
			visited[pattern] = visits + 1
			p0, p1 = movePositionToNextPattern(p0, p1, currentNumber, lenCurrentPattern)
		}
	}

	return visited[pattern] > 1
}

func movePositionToNextPattern(p0, p1 int, currentNumber string, lenCurrentPattern int) (int, int) {
	p0 = p0 + lenCurrentPattern
	p1 = p1 + lenCurrentPattern
	if p0 < len(currentNumber) && p1 > len(currentNumber) {
		p0, p1 = restartPatternPosition(p0)
	}
	return p0, p1
}

func restartPatternPosition(p0 int) (int, int) {
	return 0, p0
}
