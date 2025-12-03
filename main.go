package main

import (
	"advent-of-code-2025/common"
	"advent-of-code-2025/day1"
	"advent-of-code-2025/day2"
	"fmt"
)

func main() {
	// runDay1()
	runDay2()
}

func runDay1() {
	input := common.ParseCSV("day1/input.csv")

	// result := day1.CountTimesPointing0(50, input)

	result := day1.CountTimesHoveringAt0(50, input)

	fmt.Println("Result: ", result)
}

func runDay2() {
	input := common.ParseCSV("day2/input.csv")

	result := day2.AddingUpInvalidIds(input[0])

	fmt.Println("Result: ", result)
}
