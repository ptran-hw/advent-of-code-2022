package day2

import (
	"bufio"
	"github.com/agrison/go-commons-lang/stringUtils"
	"os"
)

const inputFile = "./day2/input.txt"

func readMatchesFromFileInput() [][]string {
	const delimiter = " "

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 3 {
			panic("input file invalid format, line length not 3")
		}

		match := []string{
			stringUtils.SubstringBefore(line, delimiter),
			stringUtils.SubstringAfter(line, delimiter),
		}
		result = append(result, match)
	}

	return result
}

func getSampleInput() [][]string {
	return [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
}
