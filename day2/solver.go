package day2

import (
	"github.com/agrison/go-commons-lang/stringUtils"
	"log"
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

type Solver struct {
}

func (s Solver) Solve() {
	//matches := getSampleMatches()
	matches := readMatchesFromFile()

	log.Printf("Using the hand guide, total score is: %d\n", calculateTotalScoreWithHandSign(matches))
	log.Printf("Using the match outcome guide, total score is: %d\n",
		calculateTotalScoreWithMatchOutcome(matches))
}

/*
Given [][]string matches, where matches[i] is a pair of values: opponent hand, my hand
Calculate the total score
*/
func calculateTotalScoreWithHandSign(matches [][]string) int {
	totalScore := 0
	for _, match := range matches {
		otherHand := parseHand(match[0])
		ownHand := parseHand(match[1])

		totalScore += calculateScore(ownHand, otherHand)
	}

	return totalScore
}

func parseHand(char string) string {
	var result string
	switch {
	case stringUtils.ContainsAnyCharacter(rockSymbols, char):
		result = rockHand
	case stringUtils.ContainsAnyCharacter(paperSymbols, char):
		result = paperHand
	case stringUtils.ContainsAnyCharacter(scissorSymbols, char):
		result = scissorHand
	default:
		log.Panicf("input contains invalid hand symbol: %s", char)
	}

	return result
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
Given [][]string matches, where matches[i] is a pair of values: opponent hand, match outcome
Calculate the total score
*/
func calculateTotalScoreWithMatchOutcome(matches [][]string) int {
	totalScore := 0
	for _, match := range matches {
		otherHand := parseHand(match[0])

		outcome := parseOutcome(match[1])
		ownHand := rotateHand(otherHand, outcome)

		totalScore += calculateScore(ownHand, otherHand)
	}

	return totalScore
}

func parseOutcome(char string) int {
	var result int
	switch char {
	case winSymbol:
		result = 1
	case drawSymbol:
		result = 0
	case loseSymbol:
		result = -1
	default:
		log.Panicf("input contains invalid outcome symbol: %s", char)
	}

	return result
}

// rock -> paper -> scissor
func rotateHand(hand string, rotations int) string {
	if rotations % 3 == 0 {
		return hand
	}

	var nextHand string
	switch hand {
	case rockHand:
		nextHand = paperHand
	case paperHand:
		nextHand = scissorHand
	case scissorHand:
		nextHand = rockHand
	}

	return rotateHand(nextHand, rotations - 1)
}
