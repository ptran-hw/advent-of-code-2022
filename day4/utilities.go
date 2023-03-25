package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./day4/input.txt"

func readAssignmentsFromFile() [][]int {
	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	result := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			panic("invalid input file format, line must contain two assignments")
		}

		assignmentA := parseAssignment(pair[0])
		assignmentB := parseAssignment(pair[1])
		result = append(result, assignmentA, assignmentB)
	}

	return result
}

func parseAssignment(entry string) []int {
	pair := strings.Split(entry, "-")
	if len(pair) != 2 {
		panic("invalid input file format, assignment must have format: {A}-{B}")
	}

	start, err := strconv.Atoi(pair[0])
	if err != nil {
		panic(fmt.Sprintf("unable to read section number: %s", pair[0]))
	}

	end, err := strconv.Atoi(pair[1])
	if err != nil {
		panic(fmt.Sprintf("unable to read section number: %s", pair[1]))
	}

	return []int{start, end}
}

func getSampleAssignments() [][]int {
	return [][]int{
		{2, 4},
		{6, 8},
		{2, 3},
		{4, 5},
		{5, 7},
		{7, 9},
		{2, 8},
		{3, 7},
		{6, 6},
		{4, 6},
		{2, 6},
		{4, 8},
	}
}
