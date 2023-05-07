package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const instructionDataFile = "./day10/instructionData.txt"
const sampleInstructionDataFile = "./day10/sampleInstructionData.txt"

func readInstructionsFromFile() []Instruction {
	return readInstructionsFromFileHelper(instructionDataFile)
}

func readSampleInstructionsFromFile() []Instruction {
	return readInstructionsFromFileHelper(sampleInstructionDataFile)
}

func readInstructionsFromFileHelper(filePath string) []Instruction {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	instructions := make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "noop" {
			instructions = append(instructions, Instruction{action: "noop"})
			continue
		}

		pair := strings.Split(line, " ")
		if len(pair) != 2 {
			panic(fmt.Sprintf("invalid input file, line must contain format: {action} {value}. line: %s", line))
		}

		action := pair[0]
		value, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}

		instructions = append(instructions, Instruction{action: action, value: value})
	}

	return instructions
}