package day2

import (
	"bufio"
	"log"
	"os"
	"strings"
)

const matchesFile = "./day2/matchesData.txt"
const matchesDelimiter = " "

func readMatchesFromFile() [][]string {
	file, err := os.Open(matchesFile)
	if err != nil {
		log.Panicf("unable to read input file: %v", err)
	}

	result := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		match := strings.Split(line, matchesDelimiter)
		if len(match) != 2 {
			log.Panicf("invalid line format: %s", line)
		}

		result = append(result, match)
	}

	return result
}

func getSampleMatches() [][]string {
	return [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
}
