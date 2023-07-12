package day5

import (
	"log"
)

type Solver struct {
}

type moveCargoWrapperFunc func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int)

func (s Solver) Solve() {
	//cargoState := getSampleInitialState()
	//instructions := getSampleInstructions()
	cargoState := readInitialStateFromFile()
	instructions := readInstructionsFromFile()

	var cargoStateCopy = deepCopy(cargoState)
	rearrangeCargoCrateByCrate(cargoStateCopy, instructions)
	printTopCrates(cargoStateCopy)

	cargoStateCopy = deepCopy(cargoState)
	rearrangeCargoWithSingleMove(cargoStateCopy, instructions)
	printTopCrates(cargoStateCopy)
}

/*
Given [][]string cargoState, and [][]int instructions,
Apply the instructions, moving cargo one at a time, and print the top cargo of each cargo stack
*/
func rearrangeCargoCrateByCrate(cargoState [][]string, instructions [][]int) {
	moveCargoFunc := func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int) {
		for cargoMoved := 0; cargoMoved < count; cargoMoved++ {
			moveCargoHelper(cargoState, cargoStartIndex, cargoEndIndex, 1)
		}
	}

	rearrangeCargo(cargoState, instructions, moveCargoFunc)
}

/*
Given [][]string cargoState, and [][]int instructions,
Apply the instructions, moving multiple cargo at the same time, and print the top cargo of each cargo stack
*/
func rearrangeCargoWithSingleMove(cargoState [][]string, instructions [][]int) {
	moveCargoFunc := func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int) {
		moveCargoHelper(cargoState, cargoStartIndex, cargoEndIndex, count)
	}

	rearrangeCargo(cargoState, instructions, moveCargoFunc)
}

func rearrangeCargo(cargoState [][]string, instructions [][]int, moveCargoFunc moveCargoWrapperFunc) {
	for _, instruction := range instructions {
		count := instruction[0]
		cargoStartIndex := instruction[1] - 1
		cargoEndIndex := instruction[2] - 1

		moveCargoFunc(cargoState, count, cargoStartIndex, cargoEndIndex)
	}
}

func moveCargoHelper(cargos [][]string, fromCargoIndex, toCargoIndex, crateCount int) {
	fromCargo := cargos[fromCargoIndex]
	toCargo := cargos[toCargoIndex]

	cutoffIndex := len(fromCargo) - crateCount
	crates := fromCargo[cutoffIndex:]
	fromCargo = fromCargo[:cutoffIndex]
	toCargo = append(toCargo, crates...)

	cargos[fromCargoIndex] = fromCargo
	cargos[toCargoIndex] = toCargo
}

func printTopCrates(cargos [][]string) {
	for index := 0; index < len(cargos); index++ {
		currCargo := cargos[index]
		if len(currCargo) > 0 {
			log.Print(currCargo[len(currCargo)-1])
		}
	}

	log.Println()
}
