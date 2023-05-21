package day12

import "fmt"

const startChar string = "S"
const endChar string = "E"
const visitedChar string = "V"
const startHeightChar string = "a"
const endHeightChar string = "z"

type Position struct {
	row int
	col int
}

type Solver struct {
}

var increments = []Position{
	{row: -1, col: 0},
	{row: 1, col: 0},
	{row: 0, col: -1},
	{row: 0, col: 1},
}

func (s Solver) Solve() {
	//grid := getSampleGrid()
	grid := readGridFromFile()

	solveShortestPathToEndPosition(grid)
	solveShortestPathFromAnyStartPositionToEndPosition(grid)
}

func solveShortestPathToEndPosition(grid [][]string) {
	duplicate := duplicateGrid(grid)

	startPosition, err := getStartPosition(duplicate)
	if err != nil {
		panic(err)
	}

	fmt.Printf("shortest path to end requires: %d steps\n", countShortestPathLengthToEndPosition(*startPosition, duplicate))
}

func solveShortestPathFromAnyStartPositionToEndPosition(grid [][]string) {
	duplicate := duplicateGrid(grid)

	startPositions := getPotentialStartPositions(duplicate)

	minSteps := -1
	for _, startPosition := range startPositions {
		steps := countShortestPathLengthToEndPosition(startPosition, duplicate)
		if minSteps == -1 {
			minSteps = steps
		}
		if steps != -1 && steps < minSteps {
			minSteps = steps
		}

		duplicate = duplicateGrid(grid)
	}

	fmt.Printf("shortest path from any start cell to end requires: %d steps\n", minSteps)
}

func duplicateGrid(grid [][]string) [][]string {
	duplicate := make([][]string, len(grid))
	for index := range grid {
		duplicate[index] = make([]string, len(grid[index]))
		copy(duplicate[index], grid[index])
	}

	return duplicate
}

func getStartPosition(grid [][]string) (*Position, error) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == startChar {
				return &Position{row: row, col: col}, nil
			}
		}
	}

	return nil, fmt.Errorf("unable to find starting position")
}

func getPotentialStartPositions(grid [][]string) []Position {
	startPositions := make([]Position, 0)
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			cell := grid[row][col]
			if  cell == startChar || cell == startHeightChar {
				startPositions = append(startPositions, Position{row: row, col: col})
			}
		}
	}

	return startPositions
}

func countShortestPathLengthToEndPosition(startPosition Position, grid [][]string) int {
	queue := make([]Position, 0)
	queue = append(queue, startPosition)

	stepsTaken := 0
	for {
		buffer := make([]Position, 0)

		for len(queue) > 0 {
			currPosition := queue[0]
			currRow, currCol := currPosition.row, currPosition.col
			queue = queue[1:]

			if grid[currRow][currCol] == endChar {
				return stepsTaken
			}

			currValue := grid[currRow][currCol]
			for _, increment := range increments {
				nextRow, nextCol := currRow + increment.row, currCol + increment.col

				if nextRow >= 0 && nextRow < len(grid) &&
					nextCol >= 0 && nextCol < len(grid[0]) &&
					canClimb(currValue, grid[nextRow][nextCol]) {
					buffer = append(buffer, Position{row: currRow + increment.row, col: currCol + increment.col})
				}
			}
			grid[currRow][currCol] = visitedChar // ensures we do not revisit cells
		}

		if len(buffer) == 0 {
			break
		} else {
			stepsTaken++
			queue = buffer
		}
	}

	return -1
}

// bug found: bytes are uint8 (0,255) so we have to be careful about arithmetic operations
func canClimb(currHeight, nextHeight string) bool {
	if nextHeight == visitedChar {
		return false
	}

	if currHeight == startChar {
		currHeight = startHeightChar
	}
	if nextHeight == endChar {
		nextHeight = endHeightChar
	}

	return nextHeight[0] <= 1 + currHeight[0]
}