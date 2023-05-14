package day11

import "fmt"

type Solver struct {
	monkeyInspectionCount []int
}

type Monkey struct {
	worryLevels []int
	operationFunc    func(int) int
	redirectTestFunc func(int) int
}

func (s *Solver) initialize(length int) {
	s.monkeyInspectionCount = make([]int, length)
}

func (s *Solver) Solve() {
	//monkeys := getSampleMonkeys()
	monkeys := readMonkeysFromFile()
	s.initialize(len(monkeys))

	for i := 1; i <= 20; i++ {
		s.processRound(monkeys)
	}

	fmt.Printf("monkey inspection count: %v\n", s.monkeyInspectionCount)
	fmt.Printf("monkey business score: %d\n", calculateMonkeyBusinessScore(s.monkeyInspectionCount))
}

func (s *Solver) processRound(monkeys []*Monkey) {
	for index := 0; index < len(monkeys); index++ {
		currMonkey := monkeys[index]
		for _, level := range currMonkey.worryLevels {
			nextWorryLevel := calculateNextWorryLevel(level, *currMonkey)
			nextMonkeyIndex := calculateNextMonkey(nextWorryLevel, *currMonkey)

			nextMonkey := monkeys[nextMonkeyIndex]
			nextMonkey.worryLevels = append(nextMonkey.worryLevels, nextWorryLevel)
		}

		s.monkeyInspectionCount[index] += len(currMonkey.worryLevels)
		currMonkey.worryLevels = make([]int, 0)
	}
}

func calculateNextWorryLevel(worryLevel int, monkey Monkey) int {
	return monkey.operationFunc(worryLevel) / 3
}

func calculateNextMonkey(nextWorryLevel int, monkey Monkey) int {
	return monkey.redirectTestFunc(nextWorryLevel)
}

func calculateMonkeyBusinessScore(inspectionCounter []int) int {
	maxScore := 0
	for i := 0; i < len(inspectionCounter) - 1; i++ {
		for j := i + 1; j < len(inspectionCounter); j++ {
			if maxScore < inspectionCounter[i] * inspectionCounter[j] {
				maxScore = inspectionCounter[i] * inspectionCounter[j]
			}
		}
	}

	return maxScore
}