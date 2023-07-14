package day8

import (
	"log"
)

type Solver struct {
}

func (s Solver) Solve() {
	//treeHeights := getTreeHeightGrid()
	treeHeights := readTreeHeightGridFromFile()

	solveVisibleTreesFromOutsideGrid(treeHeights)
	solveMaxTreeScenicScore(treeHeights)
}

/*
Given treeHeights [][]int,
Find the number of trees that is visible from outside the grid (either left, right, top, down)

A tree is visible if all trees in a straight line towards to the boundary are strictly shorter
*/
func solveVisibleTreesFromOutsideGrid(treeHeights [][]int) {
	visiblePositions := make(map[string]bool, 0)
	updateVisiblePositions(treeHeights, visiblePositions)
	log.Printf("Visible tree count: %d\n", countTotalVisibleTrees(visiblePositions))
}

/*
Given treeHeights [][]int,
Find the maximum scenic score in the grid

The scenic score of a tree is the product of viewing distances (number of trees visible) in the 4 directions
*/
func solveMaxTreeScenicScore(treeHeights [][]int) {
	log.Printf("Max scenic score: %d\n", calculateMaxScenicScore(treeHeights))
}

func updateVisiblePositions(treeHeights [][]int, visiblePositions map[string]bool) {
	const positiveStep = 1
	const negativeStep = -1

	for rowIndex := 0; rowIndex < len(treeHeights); rowIndex++ {
		updateVisibleTreeGridByRow(treeHeights, rowIndex, positiveStep, visiblePositions)
		updateVisibleTreeGridByRow(treeHeights, rowIndex, negativeStep, visiblePositions)
	}

	for colIndex := 0; colIndex < len(treeHeights[0]); colIndex++ {
		updateVisibleTreeGridByCol(treeHeights, colIndex, positiveStep, visiblePositions)
		updateVisibleTreeGridByCol(treeHeights, colIndex, negativeStep, visiblePositions)
	}
}

func updateVisibleTreeGridByRow(treeHeights [][]int, rowIndex int, increment int, visiblePositions map[string]bool) {
	row := treeHeights[rowIndex]

	var start int
	if increment > 0 {
		start = 0
	} else {
		start = len(row) - 1
	}

	maxHeight := row[start]
	for colIndex := start; isWithinBounds(colIndex, 0, len(row) - 1); colIndex += increment {
		if colIndex == 0 || colIndex == len(row) - 1 {
			visiblePositions[getKey(rowIndex, colIndex)] = true
			continue
		}

		height := row[colIndex]
		if maxHeight < height {
			visiblePositions[getKey(rowIndex, colIndex)] = true
			maxHeight = height
		}
	}
}

func updateVisibleTreeGridByCol(treeHeights [][]int, colIndex int, increment int, visiblePositions map[string]bool) {
	var start int
	if increment > 0 {
		start = 0
	} else {
		start = len(treeHeights) - 1
	}

	maxHeight := treeHeights[start][colIndex]
	for rowIndex := start; isWithinBounds(rowIndex, 0, len(treeHeights) - 1); rowIndex += increment {
		if rowIndex == 0 || rowIndex == len(treeHeights) - 1 {
			visiblePositions[getKey(rowIndex, colIndex)] = true
			continue
		}

		height := treeHeights[rowIndex][colIndex]
		if maxHeight < height {
			visiblePositions[getKey(rowIndex, colIndex)] = true
			maxHeight = height
		}
	}
}

func countTotalVisibleTrees(visiblePositions map[string]bool) int {
	return len(visiblePositions)
}

func calculateMaxScenicScore(treeHeights [][]int) int {
	directionIncrements := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	maxScore := 0
	for row := 0; row < len(treeHeights); row++ {
		for col := 0; col < len(treeHeights[0]); col++ {
			score := calculateMaxScenicScoreHelper(treeHeights, directionIncrements, row, col)
			if maxScore < score {
				maxScore = score
			}
		}
	}

	return maxScore
}

func calculateMaxScenicScoreHelper(treeHeights [][]int, directionIncrements [][]int, row int, col int) int {
	score := 1

	for _, increments := range directionIncrements {
		rowIncrement := increments[0]
		colIncrement := increments[1]

		score *= calculateScenicScore(treeHeights, row, col, rowIncrement, colIncrement)
	}

	return score
}

func calculateScenicScore(treeHeights [][]int, row, col, rowIncrement, colIncrement int) int {
	startHeight := treeHeights[row][col]
	score := 0

	rowRunner := row + rowIncrement
	colRunner := col + colIncrement
	maxRow := len(treeHeights) - 1
	maxCol := len(treeHeights[row]) - 1

	for isWithinBounds(rowRunner, 0, maxRow) && isWithinBounds(colRunner, 0, maxCol) {
		score++

		currHeight := treeHeights[rowRunner][colRunner]
		if currHeight >= startHeight {
			break
		}

		rowRunner += rowIncrement
		colRunner += colIncrement
	}

	return score
}
