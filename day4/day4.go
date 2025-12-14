package day4

type position struct {
	row    int
	column int
}

func CountAccessibleRolls(rollsGrid []string) int64 {
	accesibleRolls := int64(0)
	gridDeepth := len(rollsGrid[0])
	gridLength := len(rollsGrid)

	for i := 0; i < gridLength; i++ {
		for j := 0; j < gridDeepth; j++ {
			currentPosition := position{i, j}

			if string(rollsGrid[i][j]) == "@" && isAccesible(rollsGrid, currentPosition, gridDeepth, gridLength) {
				accesibleRolls++
			}
		}
	}

	return accesibleRolls
}

func MaxRollsCanBeRemoved(rollsGrid []string) int64 {
	accesibleRollsCount := int64(0)
	gridDeepth := len(rollsGrid[0])
	gridLength := len(rollsGrid)

	for i := 0; i < gridLength; i++ {
		for j := 0; j < gridDeepth; j++ {
			currentPosition := position{i, j}

			if string(rollsGrid[currentPosition.row][currentPosition.column]) == "@" && isAccesible(rollsGrid, currentPosition, gridDeepth, gridLength) {
				accesibleRollsCount++
				currentElem := rollsGrid[currentPosition.row]
				rollsGrid[currentPosition.row] = currentElem[:currentPosition.column] + "X" + currentElem[currentPosition.column+1:]
			}
		}
	}

	if accesibleRollsCount == 0 {
		return 0
	}

	return accesibleRollsCount + MaxRollsCanBeRemoved(rollsGrid)
}

func calculateAdjacentPositions(currentPosition position, gridLength, gridDeepth int) []position {
	positions := []position{}

	for i := currentPosition.row - 1; i <= currentPosition.row+1; i++ {
		if i >= 0 && i < gridLength {
			for j := currentPosition.column - 1; j <= currentPosition.column+1; j++ {
				if j >= 0 && j < gridDeepth {
					if i != currentPosition.row || j != currentPosition.column {
						positions = append(positions, position{i, j})
					}
				}
			}
		}
	}

	return positions
}

func isAccesible(rollsGrid []string, position position, gridLength int, gridDeepth int) bool {
	adjacentPositions := calculateAdjacentPositions(position, gridLength, gridDeepth)
	countRolls := 0
	for _, ap := range adjacentPositions {
		if string(rollsGrid[ap.row][ap.column]) == "@" {
			countRolls++
		}
	}

	return countRolls < 4
}
