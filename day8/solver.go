package day8

import (
	"fmt"
)

type Solver struct {
	visiblePositions map[string]bool
}

func (s Solver) Solve() {
	s.visiblePositions = make(map[string]bool, 0)
	s.countVisibleTrees()
	calculateMaxScenicScore()
}

func (s Solver) countVisibleTrees() {
	//treeHeights := getTreeHeightGrid()
	treeHeights := readTreeHeightGridFromFile()

	for rowIndex := 0; rowIndex < len(treeHeights); rowIndex++ {
		s.updateVisibleTreeGridByRow(treeHeights, rowIndex)
	}

	for colIndex := 0; colIndex < len(treeHeights[0]); colIndex++ {
		s.updateVisibleTreeGridByCol(treeHeights, colIndex)
	}

	fmt.Printf("Visible tree count: %d\n", s.countTotalVisibleTrees())
}

func calculateMaxScenicScore() {
	treeHeights := getTreeHeightGrid()
	//treeHeights := readTreeHeightGridFromFile()

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

func (s Solver) updateVisibleTreeGridByRow(treeHeights [][]int, rowIndex int) {
	row := treeHeights[rowIndex]

	maxHeight := row[0]
	for colIndex := 0; colIndex < len(row); colIndex++ {
		if colIndex == 0 {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			continue
		}

		height := row[colIndex]
		if maxHeight < height {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			maxHeight = height
		}
	}

	maxHeightRev := row[len(row)-1]
	for colIndex := len(row) - 1; colIndex > 0; colIndex-- {
		if colIndex == len(row) - 1 {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			continue
		}

		height := row[colIndex]
		if maxHeightRev < height {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			maxHeightRev = height
		}
	}
}

func (s Solver) updateVisibleTreeGridByCol(treeHeights [][]int, colIndex int) {
	maxHeight := treeHeights[0][colIndex]
	for rowIndex := 0; rowIndex < len(treeHeights); rowIndex++ {
		if rowIndex == 0 {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			continue
		}

		height := treeHeights[rowIndex][colIndex]
		if maxHeight < height {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			maxHeight = height
		}
	}

	maxHeightRev := treeHeights[len(treeHeights)-1][colIndex]
	for rowIndex := len(treeHeights) - 1; rowIndex > 0; rowIndex-- {
		if rowIndex == len(treeHeights) - 1 {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			continue
		}

		height := treeHeights[rowIndex][colIndex]
		if maxHeightRev < height {
			s.visiblePositions[convertToKey(rowIndex, colIndex)] = true
			maxHeightRev = height
		}
	}
}

func (s Solver) countTotalVisibleTrees() int {
	return len(s.visiblePositions)
}

func convertToKey(row, col int) string {
	// provides unique key for map
	return fmt.Sprintf("x%dy%d", row, col)
}