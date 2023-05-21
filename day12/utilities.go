package day12

import (
	"bufio"
	"os"
	"strings"
)

const gridFile = "./day12/grid.txt"

func readGridFromFile() [][]string {
	file, err := os.Open(gridFile)
	if err != nil {
		panic(err)
	}

	grid := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cells := strings.Split(line, "")

		grid = append(grid, cells)
	}

	return grid
}

func getSampleGrid() [][]string {
	return [][]string{
		strings.Split("Sabqponm", ""),
		strings.Split("abcryxxl", ""),
		strings.Split("accszExk", ""),
		strings.Split("acctuvwj", ""),
		strings.Split("abdefghi", ""),
	}
}
