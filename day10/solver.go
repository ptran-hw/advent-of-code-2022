package day10

import (
	"log"
)

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

const noopAction = "noop"
const addAction = "addx"

func (s Solver) Solve() {
	instructions := readSampleInstructionsFromFile()
	//instructions := readInstructionsFromFile()

	//solveSignalStrengthReport(instructions)
	solveSpriteCRTDisplay(instructions)
}

/*
Given instructions []Instruction,
Process the instructions and keep track of register value X,
Calculate the signal strengths reported at cycles 20, 40, 60, etc and sum them together
Note: Signal strength is calculated by multiplying the register value X with cycle value

Instructions have actions: noop, addx
- noop takes one full cycle to complete
- addx takes two full cycles to complete and increments register X by instruction value
*/
func solveSignalStrengthReport(instructions []Instruction) {
	registerValues := make([]int, 0)

	registerValues = appendRegisterValues(registerValues, instructions)
	totalSignalStrength := sumSignalStrengths(registerValues)

	log.Println("Total signal strength:", totalSignalStrength)
}

/*
Given instructions []Instruction,
Process the instructions and keep track of register value X,
Display the CRT image after processing all instructions

The sprite is length 3, where the middle position corresponds to the register value X
Each cycle imprints . when there is no overlap with sprite and # when there is overlap
CRT dimensions are 40 (cols) by 6 (rows)
*/
func solveSpriteCRTDisplay(instructions []Instruction) {
	crtScreen := make([]string, 0)

	crtScreen = updateCRTScreen(crtScreen, instructions)
	print(crtScreen)
}

func appendRegisterValues(registerValues []int, instructions []Instruction) []int {
	cycleNumber := 1
	currRegister := 1

	for _, instruction := range instructions {
		var increments int
		switch instruction.action {
		case noopAction:
			increments = 1
		case addAction:
			increments = 2
		}

		for increments > 0 {
			registerValues = append(registerValues, currRegister)
			cycleNumber++
			increments--
		}

		if instruction.action == addAction {
			currRegister += instruction.value
		}
	}

	return registerValues
}

func sumSignalStrengths(registerValues []int) int {
	total := 0

	// cycle is 1-index
	for cycle := 1; cycle <= len(registerValues); cycle++ {
		if isReportCycle(cycle) {
			total += calculateSignalStrength(cycle, registerValues[cycle - 1]) // register values is 0-index
		}
	}

	return total
}

func isReportCycle(cycle int) bool {
	return (cycle - reportStartCycle) % reportIncrement == 0
}

func calculateSignalStrength(cycle, registerValue int) int {
	return cycle * registerValue
}

func updateCRTScreen(crtScreen []string, instructions []Instruction) []string {
	cycleNumber := 1
	registerValue := 1

	buffer := ""
	for _, instruction := range instructions {
		var increments int
		switch instruction.action {
		case noopAction:
			increments = 1
		case addAction:
			increments = 2
		}

		for increments > 0 {
			buffer = addPixel(cycleNumber, registerValue, buffer)
			if isRowComplete(buffer) {
				crtScreen = append(crtScreen, buffer)
				buffer = ""
			}

			cycleNumber++
			increments--
		}

		if instruction.action == addAction {
			registerValue += instruction.value
		}
	}

	return crtScreen
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

func isRowComplete(buffer string) bool {
	return len(buffer) % 40 == 0
}

func print(crtScreen []string) {
	for _, row := range crtScreen {
		log.Println(row)
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