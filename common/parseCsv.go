package common

import (
	"bufio"
	"log"
	"os"
)

func ParseCSV(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		// Using log.Fatalf will exit the program if the file can't be opened.
		// This is often acceptable for command-line tools like for Advent of Code.
		log.Fatalf("failed to open file %s: %v", fileName, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}