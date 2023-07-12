package day3

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
	"log"
)

const lowercaseInitialPriority = 1
const uppercaseInitialPriority = 27

type Solver struct {
}

func (s Solver) Solve() {
	//rucksacks := getSampleRucksacks()
	rucksacks := readRucksacksFromFile()
	//teams := getSampleTeams()
	teams := readTeamsFromFile()

	log.Printf("total priority for rucksacks: %d\n",calculateTotalPriority(rucksacks))
	log.Printf("total priority for teams: %d\n",calculateTotalPriority(teams))
}

/*
Given [][]string rucksacks, where rucksacks[i] is pair of compartments
Find the common item in each compartment pair, and sum the priority score
*/
func calculateTotalPriority(groups [][]string) int {
	totalPriority := 0
	for _, group := range groups {
		errorItem := findCommonItem(group)
		totalPriority += calculatePriority(errorItem)
	}

	return totalPriority
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
