package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const treeHeightDataFile = "./day8/treeHeightData.txt"

func readTreeHeightGridFromFile() [][]int {
	file, err := os.Open(treeHeightDataFile)
	if err != nil {
		log.Panicf("unable to read tree height file: %v", err)
	}
	defer file.Close()

	grid := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// each tree height is a single digit
		line := strings.TrimSpace(scanner.Text())

		row := make([]int, 0)
		for _, char := range strings.Split(line, "") {
			height, err := strconv.Atoi(char)
			if err != nil {
				log.Panicf("unable to read height value: %v", err)
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

func getKey(row, col int) string {
	// provides unique key for map
	return fmt.Sprintf("x%dy%d", row, col)
}

func isWithinBounds(value int, start int, end int) bool {
	return value >= start && value <= end
}