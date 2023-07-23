package day11

import (
	"log"
)

type Solver struct {
}

type Monkey struct {
	worryLevels []int
	operationFunc    func(int) int
	redirectTestFunc func(int) int
	redirectDivisibleValue int
}

const simpleRounds int = 20
const complexRounds int = 10000

func (s Solver) Solve() {
	monkeys := getSampleMonkeys()
	//monkeys := readMonkeysFromFile()

	solveSimpleMonkeyBusinessLevel(monkeys)
	solveComplexMonkeyBusinessLevel(monkeys)
}

func solveSimpleMonkeyBusinessLevel(monkeys []*Monkey) {
	reductionFunc := func(worryValue int) int {
		return worryValue / 3
	}

	monkeyInspectionCounter := make([]int, len(monkeys))
	for i := 1; i <= simpleRounds; i++ {
		processRound(monkeys, reductionFunc, monkeyInspectionCounter)
	}

	log.Println("(simple) monkey inspection count:", monkeyInspectionCounter)
	log.Println("(simple) monkey business score:", calculateMonkeyBusinessScore(monkeyInspectionCounter))
}

func solveComplexMonkeyBusinessLevel(monkeys []*Monkey) {
	reduceFactor := getReduceFactor(monkeys)
	reductionFunc := func(worryValue int) int {
		return worryValue % reduceFactor
	}

	monkeyInspectionCounter := make([]int, len(monkeys))
	for i := 1; i <= complexRounds; i++ {
		processRound(monkeys, reductionFunc, monkeyInspectionCounter)
	}

	log.Println("(complex) monkey inspection count:", monkeyInspectionCounter)
	log.Println("(complex) monkey business score:", calculateMonkeyBusinessScore(monkeyInspectionCounter))
}

func getReduceFactor(monkeys []*Monkey) int {
	values := make([]int, len(monkeys))
	for index := 0; index < len(monkeys); index++ {
		currMonkey := monkeys[index]
		values[index] = currMonkey.redirectDivisibleValue
	}

	return calculateCommonMultiple(values)
}

func calculateCommonMultiple(values []int) int {
	result := 1
	for _, value := range values {
		if value == 0 {
			continue
		}

		result *= value
	}

	return result
}

func processRound(monkeys []*Monkey, reductionFunc func(int) int, monkeyInspectionCount []int) {
	for index, currMonkey := range monkeys {
		for _, level := range currMonkey.worryLevels {
			nextWorryLevel := calculateNextWorryLevel(level, currMonkey.operationFunc, reductionFunc)
			nextMonkeyIndex := calculateNextMonkey(nextWorryLevel, currMonkey.redirectTestFunc)

			nextMonkey := monkeys[nextMonkeyIndex]
			nextMonkey.worryLevels = append(nextMonkey.worryLevels, nextWorryLevel)
		}

		monkeyInspectionCount[index] += len(currMonkey.worryLevels)
		currMonkey.worryLevels = make([]int, 0)
	}
}

func calculateNextWorryLevel(worryLevel int, operationFunc, reductionFunc func(int) int) int {
	return reductionFunc(operationFunc(worryLevel))
}

func calculateNextMonkey(nextWorryLevel int, redirectTestFunc func(int) int) int {
	return redirectTestFunc(nextWorryLevel)
}

func calculateMonkeyBusinessScore(inspectionCounter []int) int64 {
	maxScore := int64(0)
	for i := 0; i < len(inspectionCounter) - 1; i++ {
		for j := i + 1; j < len(inspectionCounter); j++ {
			currScore := int64(inspectionCounter[i]) * int64(inspectionCounter[j])
			if maxScore < currScore {
				maxScore = currScore
			}
		}
	}

	return maxScore
}