package day3

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
)

const lowercasePriorityOffset = 1
const uppercasePriorityOffset = 27

type Solver struct {
}

func (s Solver) Solve() {
	calculateTotalPriorityForRucksacks()
	calculateTotalPriorityForTeams()
}

func calculateTotalPriorityForRucksacks() {
	rucksacks := readRucksacksFromFile()

	totalPriority := 0
	for _, rucksack := range rucksacks {
		compartmentA := rucksack[0]
		compartmentB := rucksack[1]

		errorItem := findErrorItem(compartmentA, compartmentB)
		totalPriority += calculatePriority(errorItem)
	}

	fmt.Printf("total priority: %d\n", totalPriority)
}

func findErrorItem(compartmentA, compartmentB string) string {
	itemMap := make(map[string]bool)

	for _, runeValue := range compartmentA {
		itemMap[string(runeValue)] = true
	}

	for _, runeValue := range compartmentB {
		if _, found := itemMap[string(runeValue)]; found {
			return string(runeValue)
		}
	}

	panic(fmt.Sprintf("Unable to find an error item for: %s, and %s", compartmentA, compartmentB))
}

func calculateTotalPriorityForTeams() {
	teams := readTeamsFromFile()

	totalPriority := 0
	for _, team := range teams {
		rucksackA := team[0]
		rucksackB := team[1]
		rucksackC := team[2]

		badgeItem := findBadgeItem(rucksackA, rucksackB, rucksackC)
		fmt.Printf("rucksacks: %s, %s, %s and badgeItem: %s\n", rucksackA, rucksackB, rucksackC, badgeItem)
		totalPriority += calculatePriority(badgeItem)
	}

	fmt.Printf("total priority: %d\n", totalPriority)
}

func findBadgeItem(rucksackA, rucksackB, rucksackC string) string {
	itemMap := make(map[string]int)

	for _, runeValue := range rucksackA {
		itemMap[string(runeValue)] = 1
	}

	for _, runeValue := range rucksackB {
		if _, found := itemMap[string(runeValue)]; found {
			itemMap[string(runeValue)] = 2
		}
	}

	for _, runeValue := range rucksackC {
		if value, _ := itemMap[string(runeValue)]; value == 2 {
			return string(runeValue)
		}
	}

	panic(fmt.Sprintf("Unable to find an error item for: %s, %s, and %s", rucksackA, rucksackB, rucksackC))
}

func calculatePriority(item string) int {
	if stringUtils.IsAllLowerCase(item) {
		return int(item[0]-"a"[0]) + lowercasePriorityOffset
	} else {
		return int(item[0]-"A"[0]) + uppercasePriorityOffset
	}
}
