package day1

import "strconv"

func CountTimesPointing0(startsAt int64, moves []string) int64 {
	currentPoint := startsAt
	timesPoiting0 := int64(0)

	for _, move := range moves {
		currentPoint = (currentPoint + parseMove(move)) % 100
		if currentPoint == 0 {
			timesPoiting0++
		}
	}

	return timesPoiting0
}

func CountTimesHoveringAt0(startsAt int64, moves []string) int64 {
	currentPoint := startsAt
	timesHoveringAt0 := int64(0)

	for _, move := range moves {
		currentPoint = (currentPoint + parseMove(move)) % 100
		
	}

	return timesHoveringAt0
}

func parseMove(move string) int64 {
	typeOfMove := move[0]
	amountMove, _ := strconv.ParseInt(move[1:], 10, 64)

	if typeOfMove == 'L' {
		return -1 * amountMove
	}

	return amountMove
}
