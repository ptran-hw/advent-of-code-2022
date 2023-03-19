package day1

import (
	"fmt"
	"sort"
)

type Solver struct {
}

// defined on Solver to use fields and expose to client code
func (s Solver) Solve() {
	calculateMostCalories()
	calculateTopThreeCalories()
}

func calculateMostCalories() {
	elfCalories := readElfCaloriesFromFile()
	consolidatedElfCalories := consolidateCalories(elfCalories)

	maxCalories := getMaxValue(consolidatedElfCalories)

	fmt.Println("Elf with most food calories has:", maxCalories)
}

func calculateTopThreeCalories() {
	elfCalories := readElfCaloriesFromFile()
	consolidatedElfCalories := consolidateCalories(elfCalories)

	sort.Ints(consolidatedElfCalories) // orders in ascending order

	totalCalories := 0
	for index := len(consolidatedElfCalories) - 1; index >= 0 && index >= len(consolidatedElfCalories)-3; index-- {
		totalCalories += consolidatedElfCalories[index]
	}

	fmt.Println("Top three elves has total calories:", totalCalories)
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
