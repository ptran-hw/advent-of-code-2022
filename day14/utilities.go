package day14

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const signalsFile string = "./day14/rockStructureSignals.txt"

func readRockStructureCoordinates() [][]Coordinate {
	file, err := os.Open(signalsFile)
	if err != nil {
		panic(err)
	}

	coordinates := make([][]Coordinate, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " -> ")

		buffer := make([]Coordinate, 0)
		for _, token := range tokens {
			subtokens := strings.Split(token, ",")
			if len(subtokens) != 2 {
				panic(fmt.Sprintf("unexpected coordinate format: %s", token))
			}

			x, err := strconv.Atoi(subtokens[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(subtokens[1])
			if err != nil {
				panic(err)
			}

			buffer = append(buffer, Coordinate{x: x, y: y})
		}

		coordinates = append(coordinates, buffer)
	}

	return coordinates
}

func getSampleRockStructureCoordinates() [][]Coordinate {
	return [][]Coordinate{
		{
			{x: 498, y: 4},
			{x: 498, y: 6},
			{x: 496, y: 6},
		},
		{
			{x: 503, y: 4},
			{x: 502, y: 4},
			{x: 502, y: 9},
			{x: 494, y: 9},
		},
	}
}
