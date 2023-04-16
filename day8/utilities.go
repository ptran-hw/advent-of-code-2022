package day8

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const treeHeightDataFile = "./day8/treeHeightData.txt"

func readTreeHeightGridFromFile() [][]int {
	file, err := os.Open(treeHeightDataFile)
	if err != nil {
		panic(err)
	}

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// each tree height is a single digit
		line := strings.TrimSpace(scanner.Text())

		row := make([]int, 0)
		for _, char := range strings.Split(line, "") {
			height, err := strconv.Atoi(char)
			if err != nil {
				panic(fmt.Sprintf("treeHeightData file contains a line with non-digit character: %s", line))
			}

			row = append(row, height)
		}

		grid = append(grid, row)
	}

	return grid
}

func getTreeHeightGrid() [][]int {
	return [][]int{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
}
