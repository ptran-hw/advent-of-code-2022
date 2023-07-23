package day10

import (
	"bufio"
	"log"
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
		log.Panicf("unable to read instructions file: %v", err)
	}
	defer file.Close()

	instructions := make([]Instruction, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == noopAction {
			instructions = append(instructions, Instruction{action: noopAction})
			continue
		}

		pair := strings.Split(line, " ")
		if len(pair) != 2 {
			log.Panicf("unable to parse instruction, invalid format: %s", line)
		}

		action := pair[0]
		value, err := strconv.Atoi(pair[1])
		if err != nil {
			log.Panicf("unable to parse value: %v", err)
		}

		instructions = append(instructions, Instruction{action: action, value: value})
	}

	return instructions
}