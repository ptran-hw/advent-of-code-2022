package day9

import (
	"log"
)

const simpleRopeLength = 2
const complexRopeLength = 10

const leftMovement = "L"
const rightMovement = "R"
const upMovement = "U"
const downMovement = "D"

type Instruction struct {
	direction string
	distance int
}

type Position struct {
	x int
	y int
}

type Solver struct {
}

func (s Solver) Solve() {
	//instructions := getSampleInstructions()
	//instructions := getLongDistanceSampleInstructions()
	instructions := readInstructionsFromFile()

	solveSimpleRopeBridge(instructions)
	solveComplexRopeBridge(instructions)
}

/*
Given instructions []Instruction,
Count the number of positions the tail visits at least once

The rope is length 2, and the tail knot will follow the head knot
When the tail knot is not in the same row/col with the head knot (after updating it's position),
then the tail know will move diagonally to keep up
*/
func solveSimpleRopeBridge(instructions []Instruction) {
	visitedTailPositions := make(map[string]bool)
	simulateMovements(instructions, simpleRopeLength, visitedTailPositions)
	log.Println("Simple rope bridge: number of visited positions:", countVisitedPositions(visitedTailPositions))
}

/*
Given instructions []Instruction,
Count the number of positions the tail visits at least once

The rope is length 10, and non-head knots follows the preceding knot using the same pattern as the simple scenario
*/
func solveComplexRopeBridge(instructions []Instruction) {
	visitedTailPositions := make(map[string]bool)
	simulateMovements(instructions, complexRopeLength, visitedTailPositions)
	log.Println("Complex rope bridge: number of visited positions:", countVisitedPositions(visitedTailPositions))
}

func simulateMovements(instructions []Instruction, knotCount int, visitedTailPositions map[string]bool) {
	knots := make([]*Position, 0)
	for count := 1; count <= knotCount; count++ {
		position := &Position{x: 0, y: 0}
		knots = append(knots, position)
	}

	addPosition(0, 0, visitedTailPositions)
	for _, instruction := range instructions {
		processInstruction(knots, instruction, visitedTailPositions)
	}
}

func processInstruction(knots []*Position, instruction Instruction, visitedTailPositions map[string]bool) {
	if instruction.distance == 0 {
		return
	}

	headKnot := knots[0]
	updatePosition(headKnot, instruction.direction)

	for index := 1; index < len(knots); index++ {
		currKnot := knots[index]
		prevKnot := knots[index - 1]

		if isOverlappingOrAdjacent(currKnot, prevKnot) {
			break
		}

		moveCloser(currKnot, prevKnot)
		if index == len(knots) - 1 {
			addPosition(currKnot.x, currKnot.y, visitedTailPositions)
		}
	}

	instruction.distance-- // pass by value results in a copy
	processInstruction(knots, instruction, visitedTailPositions)
}

func isOverlappingOrAdjacent(nodeA, nodeB *Position) bool {
	return ((0 <= nodeA.x - nodeB.x && nodeA.x - nodeB.x <= 1) ||
		(0 <= nodeB.x - nodeA.x && nodeB.x - nodeA.x <= 1)) &&
		((0 <= nodeA.y - nodeB.y && nodeA.y - nodeB.y <= 1) ||
		(0 <= nodeB.y - nodeA.y && nodeB.y - nodeA.y <= 1))
}

func updatePosition(knot *Position, direction string) {
	switch direction {
	case leftMovement:
		knot.x--
	case rightMovement:
		knot.x++
	case upMovement:
		knot.y++
	case downMovement:
		knot.y--
	}
}

func moveCloser(currKnot, prevKnot *Position) {
	yDiff := prevKnot.y - currKnot.y
	xDiff := prevKnot.x - currKnot.x

	updateYPosition := func() {
		if yDiff > 0 {
			updatePosition(currKnot, upMovement)
		} else {
			updatePosition(currKnot, downMovement)
		}
	}

	updateXPosition := func() {
		if xDiff > 0 {
			updatePosition(currKnot, rightMovement)
		} else {
			updatePosition(currKnot, leftMovement)
		}
	}

	switch {
	case currKnot.x == prevKnot.x:
		updateYPosition()
	case currKnot.y == prevKnot.y:
		updateXPosition()

	default:
		updateYPosition()
		updateXPosition()
	}
}

func addPosition(x, y int, visitedTailPositions map[string]bool) {
	visitedTailPositions[getKey(x, y)] = true
}

func countVisitedPositions(grid map[string]bool) int {
	return len(grid)
}