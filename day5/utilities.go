package day5

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const initialStateInputFile = "./day5/initialStateInput.txt"
const instructionsInputFile = "./day5/instructionsInput.txt"

func readInitialStateFromFile() [][]string {
	file, err := os.Open(initialStateInputFile)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.Split(line, "")) // split each char
	}
	
	return result
}

func readInstructionsFromFile() [][]int {
	file, err := os.Open(instructionsInputFile)
	if err != nil {
		panic(err)
	}

	result := make([][]int, 0)

	regex := regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		if len(matches) != 4 {
			panic("invalid instructions input file format, must have format: move {A} from {B} to {C}")
		}

		count, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}

		start, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(matches[3])
		if err != nil {
			panic(err)
		}

		result = append(result, []int{count, start, end})
	}

	return result
}

func getSampleInitialState() [][]string {
	return [][]string{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
}

func getSampleInstructions() [][]int {
	return [][]int{
		{1, 2, 1},
		{3, 1, 3},
		{2, 2, 1},
		{1, 1, 2},
	}
}
