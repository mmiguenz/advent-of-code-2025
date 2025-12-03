# Advent of Code 2025 in Go

Welcome to my repository for the Advent of Code 2025! This project is my personal journey through the annual coding challenges, using the Go programming language.

## What is Advent of Code?

[Advent of Code](https://adventofcode.com/) is an Advent calendar of small programming puzzles for a variety of skill sets and skill levels that can be solved in any programming language you like. People use them as a speed contest, for interview prep, company training, university coursework, practice problems, or to challenge each other.

## Project Structure

This repository is organized to keep solutions neat and tidy.

-   **`main.go`**: The entry point of the application. It's configured to run the solution for a specific day.
-   **`dayX/`**: Each day of the challenge has its own dedicated package (e.g., `day1`, `day2`).
    -   `dayX.go`: Contains the core logic to solve the puzzle for day `X`.
    -   `dayX_test.go`: Holds unit tests to verify the solution's correctness against examples.
    -   `input.csv`: The unique puzzle input for the day.
    -   `exercise.md`: A copy of the puzzle description from the Advent of Code website.
-   **`common/`**: A shared package for utility functions that can be reused across different days' solutions, like file parsers.

## How to Run

To run the solution for any given day, you will need to have Go installed on your system.

1.  Modify `main.go` to call the function for the day you want to solve.
2.  Run the program from the root directory of the project:

```bash
go run main.go
```

For example, to execute the solution for Day 1, your `main.go` would look like this:

```go
func main() {
	input := common.ParseCSV("day1/input.csv")
	result := day1.CountTimesPointing0(50, input)
	fmt.Println("Result: ", result)
}
```