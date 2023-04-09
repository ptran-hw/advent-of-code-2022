package day8

import "fmt"

type Solver struct {
}

func (s Solver) Solve() {
	countVisibleTrees()
	calculateMaxScenicScore()
}

func countVisibleTrees() {
	//treeHeights := getTreeHeightGrid()
	treeHeights := readTreeHeightGridFromFile()
	visibleTrees := makeVisibleTreeGrid(treeHeights)

	for row := 0; row < len(treeHeights); row++ {
		updateVisibleTreeGridByRow(treeHeights, visibleTrees, row)
	}

	for col := 0; col < len(treeHeights[0]); col++ {
		updateVisibleTreeGridByCol(treeHeights, visibleTrees, col)
	}

	fmt.Printf("Visible tree count: %d\n", countTotalVisibleTrees(visibleTrees))
}

func calculateMaxScenicScore() {
	//treeHeights := getTreeHeightGrid()
	treeHeights := readTreeHeightGridFromFile()

	directionIncrements := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}

	maxScore := 0
	for row := 0; row < len(treeHeights); row++ {
		for col := 0; col < len(treeHeights[0]); col++ {
			score := 1

			for _, increments := range directionIncrements {
				rowIncrement := increments[0]
				colIncrement := increments[1]

				score *= calculateScenicScore(treeHeights, row, col, rowIncrement, colIncrement)
			}

			if maxScore < score {
				maxScore = score
			}
		}
	}

	fmt.Printf("Max scenic score: %d\n", maxScore)
}

func calculateScenicScore(treeHeights [][]int, row, col, rowIncrement, colIncrement int) int {
	startHeight := treeHeights[row][col]
	score := 0

	rowRunner := row + rowIncrement
	colRunner := col + colIncrement
	for rowRunner >= 0 && rowRunner < len(treeHeights) && colRunner >= 0 && colRunner < len(treeHeights[row]) {
		currHeight := treeHeights[rowRunner][colRunner]
		if currHeight < startHeight {
			score++
		} else {
			score++
			break
		}

		rowRunner += rowIncrement
		colRunner += colIncrement
	}

	return score
}

func makeVisibleTreeGrid(grid [][]int) [][]bool {
	visibleTrees := make([][]bool, len(grid))
	for index, row := range grid {
		visibleTrees[index] = make([]bool, len(row))
	}

	return visibleTrees
}

func updateVisibleTreeGridByRow(treeHeights [][]int, visibleTrees [][]bool, rowIndex int) {
	row := treeHeights[rowIndex]

	visibleTrees[rowIndex][0] = true
	maxHeight := row[0]
	for colIndex := 1; colIndex < len(row); colIndex++ {
		height := row[colIndex]
		if maxHeight < height {
			visibleTrees[rowIndex][colIndex] = true

			maxHeight = height
		}
	}

	visibleTrees[rowIndex][len(row)-1] = true
	maxHeightRev := row[len(row)-1]
	for colIndex := len(row) - 1; colIndex > 0; colIndex-- {
		height := row[colIndex]
		if maxHeightRev < height {
			visibleTrees[rowIndex][colIndex] = true

			maxHeightRev = height
		}
	}
}

func updateVisibleTreeGridByCol(treeHeights [][]int, visibleTrees [][]bool, colIndex int) {
	visibleTrees[0][colIndex] = true
	maxHeight := treeHeights[0][colIndex]
	for rowIndex := 1; rowIndex < len(treeHeights); rowIndex++ {
		height := treeHeights[rowIndex][colIndex]
		if maxHeight < height {
			visibleTrees[rowIndex][colIndex] = true

			maxHeight = height
		}
	}

	visibleTrees[len(visibleTrees)-1][colIndex] = true
	maxHeightRev := treeHeights[len(treeHeights)-1][colIndex]
	for rowIndex := len(treeHeights) - 1; rowIndex > 0; rowIndex-- {
		height := treeHeights[rowIndex][colIndex]
		if maxHeightRev < height {
			visibleTrees[rowIndex][colIndex] = true

			maxHeightRev = height
		}
	}
}

func countTotalVisibleTrees(visibleTrees [][]bool) int {
	total := 0

	for row := 0; row < len(visibleTrees); row++ {
		for col := 0; col < len(visibleTrees[0]); col++ {
			if visibleTrees[row][col] {
				total++
			}
		}
	}

	return total
}
