package main

import (
	"advent-of-code-2025/common"
	"advent-of-code-2025/day1"
	"fmt"
)
func main() {
	input := common.ParseCSV("day1/input.csv")
	
	// result := day1.CountTimesPointing0(50, input)

	result := day1.CountTimesHoveringAt0(50, input)

	fmt.Println("Result: ", result)
}