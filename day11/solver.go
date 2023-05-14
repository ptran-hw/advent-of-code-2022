package day11

import "fmt"

type Solver struct {
	monkeyInspectionCount []int
}

type Monkey struct {
	worryLevels []int
	operationFunc    func(int) int
	redirectTestFunc func(int) int
	redirectDivisibleValue int
}

const roundsForPart1 int = 20
const roundsForPart2 int = 10000

func (s *Solver) initialize(length int) {
	s.monkeyInspectionCount = make([]int, length)
}

func (s *Solver) Solve() {
	//monkeys := getSampleMonkeys()
	monkeys := readMonkeysFromFile()
	s.initialize(len(monkeys))
	reduceFactor := getReduceFactor(monkeys)

	//for i := 1; i <= roundsForPart1; i++ {
	for i := 1; i <= roundsForPart2; i++ {
		s.processRound(monkeys, reduceFactor)
	}

	fmt.Printf("monkey inspection count: %v\n", s.monkeyInspectionCount)
	fmt.Printf("monkey business score: %d\n", calculateMonkeyBusinessScore(s.monkeyInspectionCount))
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

func (s *Solver) processRound(monkeys []*Monkey, reduceFactor int) {
	for index := 0; index < len(monkeys); index++ {
		currMonkey := monkeys[index]
		for _, level := range currMonkey.worryLevels {
			nextWorryLevel := calculateNextWorryLevel(level, *currMonkey, reduceFactor)
			nextMonkeyIndex := calculateNextMonkey(nextWorryLevel, *currMonkey)

			nextMonkey := monkeys[nextMonkeyIndex]
			nextMonkey.worryLevels = append(nextMonkey.worryLevels, nextWorryLevel)
		}

		s.monkeyInspectionCount[index] += len(currMonkey.worryLevels)
		currMonkey.worryLevels = make([]int, 0)
	}
}

func calculateNextWorryLevel(worryLevel int, monkey Monkey, reduceFactor int) int {
	//return monkey.operationFunc(worryLevel) / 3
	// part 2 does not reduce worry level anymore
	return monkey.operationFunc(worryLevel) % reduceFactor
}

func calculateNextMonkey(nextWorryLevel int, monkey Monkey) int {
	return monkey.redirectTestFunc(nextWorryLevel)
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