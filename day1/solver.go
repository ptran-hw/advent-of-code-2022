package day1

import (
	"fmt"
	"sort"
)

const InputFile = "/Users/ptran/Git/advent-of-code/day1/input.txt"

type Solver struct {
}

// defined on Solver to use fields and expose to client code
func (d Solver) Solve() {
	solveMostCalories()
	solveTopThreeCalories()
}

func solveTopThreeCalories() {
	elfCalories := ReadElfCaloriesFromInputFile()
	consolidatedElfCalories := consolidateCalories(elfCalories)

	sort.Ints(consolidatedElfCalories)

	totalCalories := 0

	index := Max(0, len(consolidatedElfCalories)-3) // three elves with top calories

	for index < len(consolidatedElfCalories) {
		totalCalories += consolidatedElfCalories[index]
		index++
	}

	fmt.Println("Top three elves has total calories:", totalCalories)
}

func solveMostCalories() {
	elfCalories := ReadElfCaloriesFromInputFile()
	consolidatedElfCalories := consolidateCalories(elfCalories)

	maxCalories := GetMaxValue(consolidatedElfCalories)

	fmt.Println("Elf with most food calories has:", maxCalories)
}

func consolidateCalories(caloriesArr [][]int) []int {
	var result []int
	for _, calories := range caloriesArr {
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
