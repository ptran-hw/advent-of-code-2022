package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/agrison/go-commons-lang/stringUtils"
)

const elfCaloriesFile = "./day1/elfCaloriesData.txt"

func readElfCaloriesFromFile() [][]int {
	file, err := os.Open(elfCaloriesFile)
	if err != nil {
		log.Panicf("unable to read input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([][]int, 0)
	buffer := make([]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if stringUtils.IsEmpty(line) {
			result = append(result, buffer)
			buffer = make([]int, 0)
		}

		value, err := strconv.Atoi(line)
		if err != nil {
			log.Panicf("unable to parse %s into int: %v", line, err)
		}

		buffer = append(buffer, value)
	}

	if len(buffer) > 0 {
		result = append(result, buffer)
	}

	return result
}

func getSampleElfCalories() [][]int {
	return [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
}
