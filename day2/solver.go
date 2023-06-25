package day2

import (
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
)

const rockSymbols = "AX"
const paperSymbols = "BY"
const scissorSymbols = "CZ"

const loseSymbol = "X"
const drawSymbol = "Y"
const winSymbol = "Z"

// there are no enums in go, so we use constant strings
// there's not much logic to encapsulate in a struct, so keeping as string
const rockHand = "ROCK"
const paperHand = "PAPER"
const scissorHand = "SCISSOR"

const loseOutcome = "LOSE"
const drawOutcome = "OUTCOME"
const winOutcome = "WIN"

type Solver struct {
}

func (s Solver) Solve() {
	calculateTotalScoreWithHandSign()
	calculateTotalScoreWithMatchOutcome()
}

/*
Given [][]string matches, where matches[i] is pair of values: opponent hand, my hand
Calculate the total score
*/
func calculateTotalScoreWithHandSign() {
	matches := readMatchesFromFile()

	totalScore := 0
	for _, match := range matches {
		if len(match) != 2 {
			panic(fmt.Sprintf("input contains invalid match: %s", match))
		}

		otherHand := parseHand(match[0])
		ownHand := parseHand(match[1])

		totalScore += calculateScore(ownHand, otherHand)
	}

	fmt.Printf("The total score following the strategy guide is: %d\n", totalScore)
}

func parseHand(char string) string {
	switch {
	case stringUtils.ContainsAnyCharacter(rockSymbols, char):
		return rockHand
	case stringUtils.ContainsAnyCharacter(paperSymbols, char):
		return paperHand
	case stringUtils.ContainsAnyCharacter(scissorSymbols, char):
		return scissorHand
	}

	panic(fmt.Sprintf("input contains invalid hand symbol: %s", char))
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

/*
Given [][]string matches, where matches[i] is pair of values: opponent hand, match outcome
Calculate the total score
*/
func calculateTotalScoreWithMatchOutcome() {
	matches := readMatchesFromFile()

	totalScore := 0
	for _, match := range matches {
		if len(match) != 2 {
			panic(fmt.Sprintf("input contains invalid match: %s", match))
		}

		otherHand := parseHand(match[0])
		outcome := parseOutcome(match[1])
		ownHand := convertToHand(otherHand, outcome)

		totalScore += calculateScore(ownHand, otherHand)
	}

	fmt.Printf("The total score following the strategy guide is: %d\n", totalScore)
}

func parseOutcome(char string) string {
	switch char {
	case winSymbol:
		return winOutcome
	case drawSymbol:
		return drawOutcome
	case loseSymbol:
		return loseOutcome
	}

	panic(fmt.Sprintf("input contains invalid outcome symbol: %s", char))
}

func convertToHand(otherHand, outcome string) string {
	switch outcome {
	case loseOutcome:
		return getLosingHand(otherHand)
	case drawOutcome:
		return otherHand
	case winOutcome:
		return getLosingHand(getLosingHand(otherHand))
	}

	panic(fmt.Sprintf("input contains invalid outcome: %s", outcome))
}

func getLosingHand(hand string) string {
	switch hand {
	case rockHand:
		return scissorHand
	case paperHand:
		return rockHand
	case scissorHand:
		return paperHand
	}

	panic(fmt.Sprintf("input contains invalid hand: %s", hand))
}
