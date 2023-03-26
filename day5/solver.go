package day5

import "fmt"

type Solver struct {
}

type moveCargoWrapperFunc func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int)

func (s Solver) Solve() {
	rearrangeCargoCrateByCrate()
	rearrangeCargoWithSingleMove()
}

func rearrangeCargoCrateByCrate() {
	rearrangeCargo(func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int) {
		for cargoMoved := 0; cargoMoved < count; cargoMoved++ {
			moveCargo(cargoState, cargoStartIndex, cargoEndIndex, 1)
		}
	})
}

func rearrangeCargoWithSingleMove() {
	rearrangeCargo(func(cargoState [][]string, count, cargoStartIndex, cargoEndIndex int) {
		moveCargo(cargoState, cargoStartIndex, cargoEndIndex, count)
	})
}

func rearrangeCargo(wrapperFunc moveCargoWrapperFunc) {
	//cargoState := getSampleInitialState()
	cargoState := readInitialStateFromFile()
	//instructions := getSampleInstructions()
	instructions := readInstructionsFromFile()

	for _, instruction := range instructions {
		count := instruction[0]
		cargoStartIndex := instruction[1] - 1
		cargoEndIndex := instruction[2] - 1

		wrapperFunc(cargoState, count, cargoStartIndex, cargoEndIndex)
	}

	printTopCrates(cargoState)
}

func moveCargo(cargos [][]string, fromCargoIndex, toCargoIndex, crateCount int) {
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
			fmt.Print(currCargo[len(currCargo)-1])
		}
	}

	fmt.Println()
}
