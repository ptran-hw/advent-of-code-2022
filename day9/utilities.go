package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const instructionDataFile = "./day9/instructionData.txt"

func readInstructionsFromFile() []Instruction {
	file, err := os.Open(instructionDataFile)
	if err != nil {
		log.Panicf("unable to read instructions file: %v", err)
	}
	defer file.Close()

	instructions := make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		details := strings.Split(line, " ")
		if len(details) != 2 {
			log.Panicf("unable to parse instruction, invalid format: %s", line)
		}

		direction := details[0]
		distance, err := strconv.Atoi(details[1])
		if err != nil {
			log.Panicf("unable to parse distance value: %v", err)
		}

		instructions = append(instructions, Instruction{direction: direction, distance: distance})
	}

	return instructions
}

func getSampleInstructions() []Instruction {
	return []Instruction {
		{direction: "R", distance: 4},
		{direction: "U", distance: 4},
		{direction: "L", distance: 3},
		{direction: "D", distance: 1},
		{direction: "R", distance: 4},
		{direction: "D", distance: 1},
		{direction: "L", distance: 5},
		{direction: "R", distance: 2},
	}
}

func getLongDistanceSampleInstructions() []Instruction {
	return []Instruction{
		{direction: "R", distance: 5},
		{direction: "U", distance: 8},
		{direction: "L", distance: 8},
		{direction: "D", distance: 3},
		{direction: "R", distance: 17},
		{direction: "D", distance: 10},
		{direction: "L", distance: 25},
		{direction: "U", distance: 20},
	}
}

func getKey(x, y int) string {
	return fmt.Sprintf("x%dy%d", x, y)
}