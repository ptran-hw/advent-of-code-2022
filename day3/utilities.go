package day3

import (
	"bufio"
	"log"
	"os"
)

const rucksacksFile = "./day3/rucksacksData.txt"

func readRucksacksFromFile() [][]string {
	file, err := os.Open(rucksacksFile)
	if err != nil {
		log.Panicf("unable to read rucksacks file: %v", err)
	}
	defer file.Close()

	result := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		midIndex := len(line)/2
		compartments := []string{
			line[0:midIndex],
			line[midIndex:],
		}
		result = append(result, compartments)
	}

	return result
}

func readTeamsFromFile() [][]string {
	file, err := os.Open(rucksacksFile)
	if err != nil {
		log.Panicf("unable to read input file: %v", err)
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
