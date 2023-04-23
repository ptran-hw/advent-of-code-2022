package day9

import (
	"fmt"
)

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
	visitedTailPositions map[string]bool
}

func (s Solver) Solve() {
	instructions := getSampleInstructions()
	//instructions := getLongDistanceSampleInstructions()
	//instructions := readInstructionsFromFile()

	s.solveSimpleRopeBridge(instructions)
	s.solveComplexRopeBridge(instructions)
}

func (s Solver) solveSimpleRopeBridge(instructions []Instruction) {
	s.visitedTailPositions = make(map[string]bool)
	s.simulateMovements(instructions, 2)
	fmt.Printf("Simple rope bridge: number of visited positions: %d\n", countVisitedPositions(s.visitedTailPositions))
}

func (s Solver) solveComplexRopeBridge(instructions []Instruction) {
	s.visitedTailPositions = make(map[string]bool)
	s.simulateMovements(instructions, 10)
	fmt.Printf("Complex rope bridge: number of visited positions: %d\n", countVisitedPositions(s.visitedTailPositions))
}

func (s Solver) simulateMovements(instructions []Instruction, knotCount int) {
	knots := make([]*Position, 0)
	for count := 1; count <= knotCount; count++ {
		position := &Position{x: 0, y: 0}
		knots = append(knots, position)
	}

	s.addPosition(0, 0)
	for _, instruction := range instructions {
		s.processComplexInstruction(knots, instruction)
	}
}

func (s Solver) processComplexInstruction(knots []*Position, instruction Instruction) {
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
			s.addPosition(currKnot.x, currKnot.y)
		}
	}

	s.processComplexInstruction(knots, Instruction{direction: instruction.direction, distance: instruction.distance - 1})
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

func (s Solver) addPosition(x, y int) {
	// provide a unique key
	s.visitedTailPositions[fmt.Sprintf("x%dy%d", x, y)] = true
}

func countVisitedPositions(grid map[string]bool) int {
	return len(grid)
}