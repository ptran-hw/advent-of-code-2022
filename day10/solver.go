package day10

import "fmt"

type Instruction struct {
	action string
	value int
}

type Solver struct {
	report    []int
	crtScreen []string
}

const reportStartCycle = 20
const reportIncrement = 40

func (s *Solver) initialize() {
	s.report = make([]int, 0)
	s.crtScreen = make([]string, 0)
}

// actions will pause for Y cycles
// - noop takes one full cycle
// - addx takes two full cycles
// signal strength calculated from cycle number (increments of 40, starting at 20) * X value
func (s *Solver) Solve() {
	//s.updateSignalStrengthReport()
	//fmt.Printf("Total signal strength: %d\n", totalSum(s.report))

	s.updateCRTScreen()
	printCRTScreen(s.crtScreen)
}

func (s *Solver) updateSignalStrengthReport() {
	instructions := readSampleInstructionsFromFile()
	//instructions := readInstructionsFromFile()

	cycleNumber := 1
	registerValue := 1

	for _, instruction := range instructions {
		if instruction.action == "noop" {
			cycleNumber = s.handleReportNoopAction(cycleNumber, registerValue)
		} else {
			cycleNumber, registerValue = s.handleReportAddxAction(cycleNumber, registerValue, instruction.value)
		}
	}
}

func (s *Solver) handleReportNoopAction(cycleNumber, registerValue int) int {
	s.maybeAddSignalStrength(cycleNumber, registerValue)

	return cycleNumber + 1
}

func (s *Solver) handleReportAddxAction(cycleNumber, registerValue, addValue int) (int, int) {
	s.maybeAddSignalStrength(cycleNumber, registerValue)
	s.maybeAddSignalStrength(cycleNumber + 1, registerValue)

	return cycleNumber + 2, registerValue + addValue
}

func (s *Solver) maybeAddSignalStrength(cycleNumber, registerValue int) {
	if isReportCycle(cycleNumber) {
		s.report = append(s.report, calculateSignalStrength(cycleNumber, registerValue))
	}
}

func isReportCycle(cycle int) bool {
	return (cycle - reportStartCycle) % reportIncrement == 0
}

func calculateSignalStrength(cycle, registerValue int) int {
	return cycle * registerValue
}

func totalSum(nums []int) int {
	total := 0
	for _, value := range nums {
		total += value
	}

	return total
}

func (s *Solver) updateCRTScreen() {
	//instructions := readSampleInstructionsFromFile()
	instructions := readInstructionsFromFile()

	cycleNumber := 1
	registerValue := 1

	buffer := ""
	for _, instruction := range instructions {
		if instruction.action == "noop" {
			cycleNumber, buffer = s.handleNoopAction(cycleNumber, registerValue, buffer)
		} else {
			cycleNumber, registerValue, buffer = s.handleAddxAction(cycleNumber, registerValue, instruction.value, buffer)
		}
	}
}

func (s *Solver) handleNoopAction(cycleNumber, registerValue int, buffer string) (int, string) {
	buffer = addPixel(cycleNumber, registerValue, buffer)
	buffer = s.maybeAddCRTScreenRow(cycleNumber, buffer)

	return cycleNumber + 1, buffer
}

func (s *Solver) handleAddxAction(cycleNumber, registerValue, addValue int, buffer string) (int, int, string) {
	buffer = addPixel(cycleNumber, registerValue, buffer)
	buffer = s.maybeAddCRTScreenRow(cycleNumber, buffer)
	buffer = addPixel(cycleNumber + 1, registerValue, buffer)
	buffer = s.maybeAddCRTScreenRow(cycleNumber + 1, buffer)

	return cycleNumber + 2, registerValue + addValue, buffer
}

func addPixel(cycleNumber, registerValue int, buffer string) string {
	if isOverlappingSprite(cycleNumber, registerValue) {
		return buffer + "#"
	} else {
		return buffer + "."
	}
}

func isOverlappingSprite(cycle, registerValue int) bool {
	position := (cycle - 1) % 40 // adjusting for 1-index offset

	return position == registerValue || position == registerValue - 1 || position == registerValue + 1
}

func (s *Solver) maybeAddCRTScreenRow(cycleNumber int, buffer string) string {
	if isRowComplete(cycleNumber) {
		s.crtScreen = append(s.crtScreen, buffer)
		buffer = ""
	}

	return buffer
}

func isRowComplete(cycle int) bool {
	return cycle % 40 == 0
}

func printCRTScreen(rows []string) {
	for _, row := range rows {
		fmt.Printf("%v\n", row)
	}
}

// Appendix
// attempt 3
//for _, instruction := range instructions {
//	if instruction.action == "noop" {
//		cycleNumber++
//		maybeAddSignalStrength(cycleNumber, registerValue)
//		continue
//	} else {
//		cycleNumber++
//		maybeAddSignalStrength(cycleNumber, registerValue)
//		cycleNumber++
//		maybeAddSignalStrength(cycleNumber, registerValue)
//
//		registerValue += instruction.value
//	}
//}

// attempt 2
//for _, instruction := range instructions {
//	maybeAddSignalStrength(cycleNumber, registerValue)
//	if instruction.action == "noop" {
//		cycleNumber++
//		continue
//	}
//
//	cycleNumber++
//	maybeAddSignalStrength(cycleNumber, registerValue)
//	cycleNumber++
//	registerValue += instruction.value
//}

// Attempt 1
//for _, instruction := range instructions {
//	// process current state
//	if isReportCycle(cycleNumber) {
//		report = append(report, calculateSignalStrength(cycleNumber, registerValue))
//	}
//	cycleNumber++
//
//	// process current instruction?
//	if instruction.action == "noop" {
//		continue
//	}
//
//	// process next state
//	if isReportCycle(cycleNumber) {
//		report = append(report, calculateSignalStrength(cycleNumber, registerValue))
//	}
//	cycleNumber++
//
//	registerValue += instruction.value
//}