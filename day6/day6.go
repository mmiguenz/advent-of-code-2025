package day6

import (
	"strconv"
	"strings"
)

func GrandTotal(homeWorkLines []string) int64 {
	problems := []*[]string{}
	for _, line := range homeWorkLines {
		elems := strings.Split(line, " ")

		for problemIndex, elm := range removeEmptyStrings(elems) {
			problem := &[]string{}
			if len(problems) > problemIndex {
				problem = problems[problemIndex]
			} else {
				problems = append(problems, problem)
			}

			*problem = append(*problem, elm)

		}
	}

	result := int64(0)
	for _, problem := range problems {
		operator, initValue := getOperator((*problem)[len(*problem)-1])
		result += foldl((*problem)[:len(*problem)-1], initValue, operator)
	}

	return result
}

func GrandTotalV2(homeWorkLines []string) int64 {
	totalRows := len(homeWorkLines)
	totalColmumns := len(homeWorkLines[0])
	currentColumn := totalColmumns - 1
	grandTotal := int64(0)
	columns := []string{}

	for currentColumn >= 0 {
		column := ""
		var operator func(acc, x int64) int64
		var initValue int64

		for row := 0; row < totalRows; row++ {
			value := string(homeWorkLines[row][currentColumn])
			if value != " " {
				if value == "*" || value == "+" {
					operator, initValue = getOperator(value)
				} else {
					column += value
				}
			}
		}

		columns = append(columns, column)
		if operator != nil {
			grandTotal += foldl(columns, initValue, operator)
			columns = []string{}
			currentColumn -= 2
		} else {
			currentColumn--
		}

	}

	return grandTotal

}

func getOperator(op string) (func(acc, x int64) int64, int64) {
	switch op {
	case "*":
		return func(acc, x int64) int64 {
			return acc * x
		}, 1
	case "+":
		return func(acc, x int64) int64 {
			return acc + x
		}, 0
	default:
		return nil, 0
	}

}

func foldl(slice []string, init int64, f func(acc, x int64) int64) int64 {
	acc := init
	for _, x := range slice {
		parsed, err := strconv.ParseInt(x, 10, 64)
		if err != nil {
			panic(err)
		}
		acc = f(acc, parsed)
	}
	return acc
}

func removeEmptyStrings(slice []string) []string {
	result := slice[:0] // use the same underlying array to avoid extra allocation
	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	return result
}
