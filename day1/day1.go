package day1

import (
	"fmt"
	"math"
	"strconv"
)

func CountTimesPointing0(startsAt int64, moves []string) int64 {
	currentPoint := startsAt
	timesPoiting0 := int64(0)

	for _, move := range moves {
		movementResult := currentPoint + parseMove(move)

		if movementResult < 0 {
			currentPoint = 100 + movementResult
		} else {
			currentPoint = movementResult % 100
		}

		if currentPoint == 0 {
			timesPoiting0++
		}
	}

	return timesPoiting0
}

func CountTimesHoveringAt0(startsAt int64, moves []string) int64 {
	currentPoint := startsAt
	timesHovering0 := int64(0)

	for _, move := range moves {
		parsedMove := parseMove(move)

		// Use the absolute value to correctly count full rotations
		timesHovering0 += int64(math.Abs(float64(parsedMove))) / 100
		parsedMove = parsedMove % 100

		movementResult := currentPoint + parsedMove

		if movementResult < 0 {
			currentPoint = movementResult + 100
			timesHovering0++
		} else {
			if movementResult >= 100 {
				timesHovering0++
				currentPoint = movementResult % 100
			} else {
				currentPoint = movementResult
			}

		}

		fmt.Println("Current Position: ", currentPoint)

	}

	return timesHovering0
}

func parseMove(move string) int64 {
	typeOfMove := move[0]
	amountMove, _ := strconv.ParseInt(move[1:], 10, 64)

	if typeOfMove == 'L' {
		return -1 * amountMove
	}

	return amountMove
}
