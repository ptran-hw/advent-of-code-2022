package day5

import "fmt"

type Solver struct {
}

func (s Solver) Solve() {
	rearrangeCargoCrateByCrate()
	rearrangeCargoWithNewCrane()
}

func rearrangeCargoCrateByCrate() {
	//cargoState := getSampleInitialState()
	cargoState := readInitialStateFromFile()
	//instructions := getSampleInstructions()
	instructions := readInstructionsFromFile()

	for _, instruction := range instructions {
		count := instruction[0]
		cargoStartIndex := instruction[1] - 1
		cargoEndIndex := instruction[2] - 1

		for cargoMoved := 0; cargoMoved < count; cargoMoved++ {
			moveCargo(cargoState, cargoStartIndex, cargoEndIndex)
		}
	}

	printTopCrates(cargoState)
}

func moveCargo(cargos [][]string, startIndex, endIndex int) {
	startCargo := cargos[startIndex]
	endCargo := cargos[endIndex]

	n := len(startCargo) - 1
	crate := startCargo[n]
	startCargo = startCargo[:n]
	endCargo = append(endCargo, crate)

	cargos[startIndex] = startCargo
	cargos[endIndex] = endCargo
}

func rearrangeCargoWithNewCrane() {
	//cargoState := getSampleInitialState()
	cargoState := readInitialStateFromFile()
	//instructions := getSampleInstructions()
	instructions := readInstructionsFromFile()

	for _, instruction := range instructions {
		count := instruction[0]
		cargoStartIndex := instruction[1] - 1
		cargoEndIndex := instruction[2] - 1

		moveMultipleCargos(cargoState, cargoStartIndex, cargoEndIndex, count)
	}

	printTopCrates(cargoState)
}

func moveMultipleCargos(cargos [][]string, fromCargoIndex, toCargoIndex, crateCount int) {
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
