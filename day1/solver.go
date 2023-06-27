package day1

import (
	"log"
	"sort"
)

type Solver struct {
}

// defined on Solver to use fields and expose to client code
func (s Solver) Solve() {
	//elfCalories := getSampleElfCalories()
	elfCalories := readElfCaloriesFromFile()

	calculateMostCalories(elfCalories)
	calculateTopThreeCalories(elfCalories)
}

/*
Given [][]int elfCalories, where elfCalories[i] is a list of calories for elf i
Find the elf with the maximum total calories and return the maximum total calories value
*/
func calculateMostCalories(elfCalories [][]int) {
	calculateTopCaloriesHelper(elfCalories, 1)
}

/*
Given [][]int elfCalories, where elfCalories[i] is a list of calories for elf i
Find the top 3 elves in terms of maximum total calories and return the sum of their total calories values
*/
func calculateTopThreeCalories(elfCalories [][]int) {
	calculateTopCaloriesHelper(elfCalories, 3)
}

func calculateTopCaloriesHelper(elfCalories [][]int, elfCount int) {
	consolidatedElfCalories := consolidateCalories(elfCalories)

	sort.Ints(consolidatedElfCalories) // orders in ascending order

	totalCalories := 0
	remaining := elfCount
	for index := len(consolidatedElfCalories) - 1; index >= 0 && index >= len(consolidatedElfCalories)-3; index-- {
		totalCalories += consolidatedElfCalories[index]
		remaining--

		if remaining == 0 {
			break
		}
	}

	log.Printf("Top %d efl/elves has total calories: %d\n", elfCount, totalCalories)
}

func consolidateCalories(elfCalories [][]int) []int {
	var result []int
	for _, calories := range elfCalories {
		totalCalories := calculateTotalCalories(calories)
		result = append(result, totalCalories)
	}

	return result
}

func calculateTotalCalories(foodList []int) int {
	total := 0
	for _, calories := range foodList {
		total += calories
	}

	return total
}
