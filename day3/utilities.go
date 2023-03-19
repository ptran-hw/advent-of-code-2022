package day3

import (
	"bufio"
	"os"
)

const inputFileA = "./day3/inputA.txt"

func readRucksacksFromFile() [][]string {
	file, err := os.Open(inputFileA)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		compartments := []string{
			line[0 : len(line)/2],
			line[len(line)/2:],
		}
		result = append(result, compartments)
	}

	return result
}

func readTeamsFromFile() [][]string {
	file, err := os.Open(inputFileA)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)
	buffer := make([]string, 0, 3)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		buffer = append(buffer, line)

		if len(buffer) == 3 {
			result = append(result, buffer)
			buffer = make([]string, 0, 3)
		}
	}

	return result
}

func getSampleRucksacks() [][]string {
	return [][]string{
		{"vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrV", "vPwwTWBwg"},
		{"wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn"},
		{"ttgJtRGJ", "QctTZtZT"},
		{"CrZsJsPPZsGz", "wwsLwLmpwMDw"},
	}
}

func getSampleTeams() [][]string {
	return [][]string{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "PmmdzqPrVvPwwTWBwg"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "ttgJtRGJQctTZtZT", "CrZsJsPPZsGzwwsLwLmpwMDw"},
	}
}
