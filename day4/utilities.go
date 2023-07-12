package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const assignmentsFile = "./day4/assignmentsData.txt"

func readAssignmentsFromFile() [][]Assignment {
	file, err := os.Open(assignmentsFile)
	if err != nil {
		log.Panicf("unable to read input file: %v", err)
	}

	result := make([][]Assignment, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		pair := strings.Split(line, ",")
		if len(pair) != 2 {
			panic("invalid input file format, line must contain two assignments")
		}

		assignmentA := parseAssignment(pair[0])
		assignmentB := parseAssignment(pair[1])
		result = append(result, []Assignment{assignmentA, assignmentB})
	}

	return result
}

func parseAssignment(entry string) Assignment {
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

	return Assignment{start: start, end: end}
}

func getSampleAssignments() [][]Assignment {
	return [][]Assignment{
		{
			{start: 2, end: 4},
			{start: 6, end: 8},
		},
		{
			{start: 2, end: 3},
			{start: 4, end: 5},
		},
		{
			{start: 5, end: 7},
			{start: 7, end: 9},
		},
		{
			{start: 2, end: 8},
			{start: 3, end: 7},
		},
		{
			{start: 6, end: 6},
			{start: 4, end: 6},
		},
		{
			{start: 2, end: 6},
			{start: 4, end: 8},
		},
	}
}
