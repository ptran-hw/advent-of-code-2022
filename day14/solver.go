package day14

import (
	"fmt"
)

type Coordinate struct {
	x int
	y int
}

type Solver struct {
}

var sandEntryCoordinate = Coordinate{x: 500, y: 0}

func (s Solver) Solve() {
	rockStructures := getSampleRockStructureCoordinates()
	//rockStructures := readRockStructureCoordinates()

	solveMaximumSandComeToRest(rockStructures)
	solveMaximumSandComeToRestWithFloor(rockStructures)
}

func solveMaximumSandComeToRest(rockStructures [][]Coordinate) {
	coordinates := getAllRockStructureCoordinates(rockStructures)

	maxY := findMaxY(coordinates)
	grid := createGridMap(coordinates)

	sandCounter := 0
	for addSand(grid, maxY) {
		sandCounter++
	}

	fmt.Printf("the maximum amount of sand added: %d\n", sandCounter)
}

func solveMaximumSandComeToRestWithFloor(rockStructures [][]Coordinate) {
	coordinates := getAllRockStructureCoordinates(rockStructures)

	maxY := findMaxY(coordinates)
	floorY := maxY + 2
	grid := createGridMap(coordinates)

	sandCounter := 0
	for addSandWithFloor(grid, floorY) {
		sandCounter++
	}

	fmt.Printf("[with floor] the maximum amount of sand added: %d\n", sandCounter)
}

func getAllRockStructureCoordinates(rockStructures [][]Coordinate) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for _, rockStructure := range rockStructures {
		for _, coordinate := range rockStructure {
			coordinates = append(coordinates, coordinate)
		}
	}

	// coordinates between start and end coordinates
	for _, rockStructure := range rockStructures {
		currCoordinate := rockStructure[0]

		for _, nextCoordinate := range rockStructure {
			if currCoordinate == nextCoordinate {
				continue
			}

			xDiff := nextCoordinate.x - currCoordinate.x
			yDiff := nextCoordinate.y - currCoordinate.y

			switch {
			case xDiff != 0:
				var startX int
				var endX int
				if currCoordinate.x <= nextCoordinate.x {
					startX = currCoordinate.x
					endX = nextCoordinate.x
				} else {
					startX = nextCoordinate.x
					endX = currCoordinate.x
				}

				for x := startX + 1; x < endX; x++ {
					coordinates = append(coordinates, Coordinate{x: x, y: currCoordinate.y})
				}
			case yDiff != 0:
				var startY int
				var endY int
				if currCoordinate.y <= nextCoordinate.y {
					startY = currCoordinate.y
					endY = nextCoordinate.y
				} else {
					startY = nextCoordinate.y
					endY = currCoordinate.y
				}

				for y := startY + 1; y < endY; y++ {
					coordinates = append(coordinates, Coordinate{x: currCoordinate.x, y: y})
				}
			}

			currCoordinate = nextCoordinate
		}
	}

	return coordinates
}

func findMaxY(coordinates []Coordinate) int {
	maxValue := 0
	for _, coordinate := range coordinates {
		if maxValue < coordinate.y {
			maxValue = coordinate.y
		}
	}

	return maxValue
}

func createGridMap(coordinates []Coordinate) map[string]bool {
	grid := make(map[string]bool, 0)

	for _, coordinate := range coordinates {
		key := getKey(coordinate.x, coordinate.y)
		grid[key] = true
	}

	return grid
}

func getKey(x, y int) string {
	return fmt.Sprintf("%d,%d", x, y)
}

func addSand(grid map[string]bool, maxY int) bool {
	return addSandHelper(sandEntryCoordinate.x, sandEntryCoordinate.y, grid, maxY)
}

func addSandHelper(x, y int, grid map[string]bool, maxY int) bool {
	if y >= maxY {
		return false
	}

	switch {
	case !grid[getKey(x, y+1)]:
		return addSandHelper(x, y+1, grid, maxY)
	case !grid[getKey(x-1, y+1)]:
		return addSandHelper(x-1, y+1, grid, maxY)
	case !grid[getKey(x+1, y+1)]:
		return addSandHelper(x+1, y+1, grid, maxY)
	default:
		grid[getKey(x, y)] = true
		return true
	}
}

func addSandWithFloor(grid map[string]bool, floorY int) bool {
	return addSandWithFloorHelper(sandEntryCoordinate.x, sandEntryCoordinate.y, grid, floorY)
}

func addSandWithFloorHelper(x, y int, grid map[string]bool, floorY int) bool {
	// changed base case
	if grid[getKey(x, y)] {
		return false
	}

	switch {
	// handle scenario with Floor
	case y + 1 == floorY:
		grid[getKey(x, y)] = true
		return true
	case !grid[getKey(x, y+1)]:
		return addSandWithFloorHelper(x, y+1, grid, floorY)
	case !grid[getKey(x-1, y+1)]:
		return addSandWithFloorHelper(x-1, y+1, grid, floorY)
	case !grid[getKey(x+1, y+1)]:
		return addSandWithFloorHelper(x+1, y+1, grid, floorY)
	default:
		grid[getKey(x, y)] = true
		return true
	}
}