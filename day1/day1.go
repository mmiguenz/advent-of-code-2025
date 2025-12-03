package day1

import (
	"fmt"
	"math"
	"strconv"
)

const dialSize = 100
func CountTimesPointing0(startsAt int64, moves []string) int64 {
	currentPoint := startsAt
	timesPoiting0 := int64(0)

	for _, move := range moves {
		movementResult := currentPoint + parseMove(move)

		if movementResult < 0 {
			currentPoint = dialSize + movementResult
		} else {
			currentPoint = movementResult % dialSize
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

		timesHovering0 += timesSpiningAndReturnedToSamePosition(parsedMove)

		parsedMove = parsedMove % dialSize

		movementResult := currentPoint + parsedMove

		if movementResult < 0 {
			if currentPoint != 0 {
				timesHovering0++
			}
			currentPoint = movementResult + dialSize
			
		} else {
			if movementResult >= dialSize {
				timesHovering0++
				currentPoint = movementResult % dialSize
			} else {
				currentPoint = movementResult
				if currentPoint == 0 {
					timesHovering0++
				}
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

func timesSpiningAndReturnedToSamePosition(parsedMove int64) int64 {
	return int64(math.Abs(float64(parsedMove))) / dialSize
}
