package day2

import (
	"bufio"
	"fmt"
	"os"

	"github.com/agrison/go-commons-lang/stringUtils"
)

type SequelSolver struct {
}

type Hand struct {
	value string
}

func NewHand(handChar string) *Hand {
	switch handChar {
	case rockSymbol:
		return &Hand{rockHand}
	case paperSymbol:
		return &Hand{paperHand}
	case scissorSymbol:
		return &Hand{scissorHand}
	default:
		return nil
	}
}

func (h *Hand) getLosingHand() *Hand {
	if h == nil {
		return nil
	}

	switch h.value {
	case rockHand:
		return &Hand{scissorHand}
	case paperHand:
		return &Hand{rockHand}
	case scissorHand:
		return &Hand{paperHand}
	}

	return nil
}

func (h *Hand) copy() *Hand {
	if h == nil {
		return nil
	}

	return &Hand{h.value}
}

func (h *Hand) getOpponentHand(outcome string) *Hand {
	if h == nil {
		return nil
	}

	switch outcome {
	case winSymbol:
		return h.getLosingHand().getLosingHand()
	case loseSymbol:
		return h.getLosingHand()
	default:
		return h.copy()
	}
}

const rockSymbol = "A"
const paperSymbol = "B"
const scissorSymbol = "C"

const loseSymbol = "X"
const drawSymbol = "Y"
const winSymbol = "Z"

func (s SequelSolver) Solve() {
	//input := s.getSampleInput()
	input := s.readFileInput()

	totalScore := 0
	for _, pair := range input {
		if len(pair) != 2 {
			panic(fmt.Sprintf("input contains invalid match: %s", pair))
		}

		otherHand := NewHand(pair[0])
		ownHand := otherHand.getOpponentHand(pair[1])

		totalScore += s.calculateScore(ownHand, otherHand)
	}

	fmt.Printf("The total score following the strategy guide is: %d\n", totalScore)
}

func (s SequelSolver) mapToHand(handChar string) string {
	switch {
	case handChar == rockSymbol:
		return rockHand
	case handChar == paperSymbol:
		return paperHand
	case handChar == scissorSymbol:
		return scissorHand
	default:
		panic(fmt.Sprintf("input contains invalid hand symbol: %s", handChar))
	}
}

func (s SequelSolver) calculateScore(ownHand, otherHand *Hand) int {
	score := 0
	switch ownHand.value {
	case rockHand:
		score += 1
	case paperHand:
		score += 2
	case scissorHand:
		score += 3
	}

	switch {
	case s.isWinningHand(ownHand.value, otherHand.value):
		score += 6
	case ownHand.value == otherHand.value: // TODO: refactor to use struct equality method
		score += 3
	}

	return score
}

func (s SequelSolver) isWinningHand(handA, handB string) bool {
	return (handA == rockHand && handB == scissorHand) ||
		(handA == paperHand && handB == rockHand) ||
		(handA == scissorHand && handB == paperHand)
}

func (s SequelSolver) getSampleInput() [][]string {
	return [][]string{
		{"A", "Y"},
		{"B", "X"},
		{"C", "Z"},
	}
}

// TODO: refactor and extract into utility file
func (s SequelSolver) readFileInput() [][]string {
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
