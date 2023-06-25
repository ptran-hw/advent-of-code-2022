package day3

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
)

const lowercaseInitialPriority = 1
const uppercaseInitialPriority = 27

type Solver struct {
}

func (s Solver) Solve() {
	calculateTotalPriorityForRucksacks()
	calculateTotalPriorityForTeams()
}

/*
Given [][]string rucksacks, where rucksacks[i] is pair of compartments
Find the common item in each compartment pair, and sum the priority score
*/
func calculateTotalPriorityForRucksacks() {
	rucksacks := readRucksacksFromFile()

	totalPriority := 0
	for _, rucksack := range rucksacks {
		compartmentA := rucksack[0]
		compartmentB := rucksack[1]

		errorItem := findCommonItem([]string{compartmentA, compartmentB})
		totalPriority += calculatePriority(errorItem)
	}

	fmt.Printf("total priority: %d\n", totalPriority)
}

/*
Given [][]string team, where teams[i] is tuple of rucksacks
Find the common badge item in each rucksack tuple, and sum the priority score
*/
func calculateTotalPriorityForTeams() {
	teams := readTeamsFromFile()

	totalPriority := 0
	for _, team := range teams {
		rucksackA := team[0]
		rucksackB := team[1]
		rucksackC := team[2]

		badgeItem := findCommonItem([]string{rucksackA, rucksackB, rucksackC})
		totalPriority += calculatePriority(badgeItem)
	}

	fmt.Printf("total priority: %d\n", totalPriority)
}

func findCommonItem(itemGroups []string) string {
	consecutiveCommonItemMap := make(map[string]int)

	// populate consecutiveCommonItemMap
	for index, itemGroup := range itemGroups {
		for _, runeValue := range itemGroup {
			char := string(runeValue)
			if index == 0 {
				consecutiveCommonItemMap[char] = 0
				continue
			}

			if lastIndexFound, _ := consecutiveCommonItemMap[char]; lastIndexFound == index-1 {
				consecutiveCommonItemMap[char] = index
			}
		}
	}

	// check for item count matches number of item groups
	for _, runeValue := range itemGroups[0] {
		char := string(runeValue)
		if lastIndexFound, _ := consecutiveCommonItemMap[char]; lastIndexFound == len(itemGroups)-1 {
			return char
		}
	}

	panic(fmt.Sprintf("common item not found for: %v", itemGroups))
}

func calculatePriority(item string) int {
	switch {
	case stringUtils.IsAllLowerCase(item):
		return int(item[0]-"a"[0]) + lowercaseInitialPriority
	case stringUtils.IsAllUpperCase(item):
		return int(item[0]-"A"[0]) + uppercaseInitialPriority
	default:
		return 0
	}
}
