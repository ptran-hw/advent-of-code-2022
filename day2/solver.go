package day2

import (
	"bufio"
	"fmt"
	"os"

	"github.com/agrison/go-commons-lang/stringUtils"
)

type Solver struct {
}

const rockSymbols = "AX"
const paperSymbols = "BY"
const scissorSymbols = "CZ"

const rockHand = "ROCK"
const paperHand = "PAPER"
const scissorHand = "SCISSOR"

func (s Solver) Solve() {
	//input := getSampleInput()
	input := readFileInput()

	totalScore := 0
	for _, pair := range input {
		if len(pair) != 2 {
			panic(fmt.Sprintf("input contains invalid match: %s", pair))
		}

		otherHand := mapToHand(pair[0])
		ownHand := mapToHand(pair[1])

		totalScore += calculateScore(ownHand, otherHand)
	}

	fmt.Printf("The total score following the strategy guide is: %d\n", totalScore)
}

func mapToHand(char string) string {
	switch {
	case stringUtils.ContainsAnyCharacter(rockSymbols, char):
		return rockHand
	case stringUtils.ContainsAnyCharacter(paperSymbols, char):
		return paperHand
	case stringUtils.ContainsAnyCharacter(scissorSymbols, char):
		return scissorHand
	default:
		panic(fmt.Sprintf("input contains invalid hand symbol: %s", char))
	}
}

func calculateScore(ownHand, otherHand string) int {
	score := 0
	switch ownHand {
	case rockHand:
		score += 1
	case paperHand:
		score += 2
	case scissorHand:
		score += 3
	}

	switch {
	case isWinningHand(ownHand, otherHand):
		score += 6
	case ownHand == otherHand:
		score += 3
	}

	return score
}

func isWinningHand(handA, handB string) bool {
	return (handA == rockHand && handB == scissorHand) ||
		(handA == paperHand && handB == rockHand) ||
		(handA == scissorHand && handB == paperHand)
}

func getSampleInput() [][]string {
	return [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
}

func readFileInput() [][]string {
	const inputFile = "/Users/ptran/Git/advent-of-code/day2/input.txt"
	const delimiter = " "

	file, err := os.Open(inputFile)
	if err != nil {
		panic(err)
	}

	result := make([][]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 3 {
			panic("input file invalid format, line length not 3")
		}

		round := []string{
			stringUtils.SubstringBefore(line, delimiter),
			stringUtils.SubstringAfter(line, delimiter),
		}
		result = append(result, round)
	}

	return result
}
