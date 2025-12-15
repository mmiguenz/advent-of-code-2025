package main

import (
	"advent-of-code-2025/common"
	"advent-of-code-2025/day1"
	"advent-of-code-2025/day2"
	"advent-of-code-2025/day3"
	"advent-of-code-2025/day4"
	"advent-of-code-2025/day5"
	"advent-of-code-2025/day6"
	"fmt"
	"strconv"
)

func main() {
	// runDay1()
	// runDay2()
	//runDay3()
	//runDay4()
	//runDay5()

	runDay6()
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

func runDay3() {
	input := common.ParseCSV("day3/input.csv")

	result := day3.SumAllMaxJoltage(input, 12)

	fmt.Println("Result: ", result)
}

func runDay4() {
	input := common.ParseCSV("day4/input.csv")

	//result := day4.CountAccessibleRolls(input)

	result := day4.MaxRollsCanBeRemoved(input)

	fmt.Println("Result: ", result)
}

func runDay5() {
	input := common.ParseCSV("day5/input.csv")
	ranges := []string{}
	ingredients := []int64{}
	endOfRanges := false
	for _, line := range input {
		if line == "" {
			endOfRanges = true
			continue
		}

		if endOfRanges {
			ingredient, _ := strconv.ParseInt(line, 10, 64)
			ingredients = append(ingredients, ingredient)
		} else {
			ranges = append(ranges, line)
		}

	}

	//result := day5.CountFreshIngredients(ranges, ingredients)
	result := day5.CountFreshIngredientsAvailable(ranges)

	fmt.Println("Result: ", result)
}

func runDay6() {
	input := common.ParseCSV("day6/input.csv")

	result := day6.GrandTotal(input)

	fmt.Println("Result: ", result)
}
