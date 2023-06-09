package day15

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const sensorDataFile = "./day15/sensorData.txt"
var sensorDataRegexp = regexp.MustCompile("Sensor at x=(-?\\d+), y=(-?\\d+): closest beacon is at x=(-?\\d+), y=(-?\\d+)")

func readSensorBeaconCoordinates() [][]Coordinate {
	file, err := os.Open(sensorDataFile)
	if err != nil {
		panic(err)
	}

	coordinatePairs := make([][]Coordinate, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !sensorDataRegexp.MatchString(line) {
			panic(fmt.Sprintf("invalid sensor data line: %s", line))
		}

		tokens := sensorDataRegexp.FindStringSubmatch(line)
		sensorX, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic(err)
		}
		sensorY, err := strconv.Atoi(tokens[2])
		if err != nil {
			panic(err)
		}
		beaconX, err := strconv.Atoi(tokens[3])
		if err != nil {
			panic(err)
		}
		beaconY, err := strconv.Atoi(tokens[4])
		if err != nil {
			panic(err)
		}

		pair := []Coordinate {
			{x: sensorX, y: sensorY},
			{x: beaconX, y: beaconY},
		}
		coordinatePairs = append(coordinatePairs, pair)
	}

	return coordinatePairs
}

func getSampleSensorBeaconCoordinates() [][]Coordinate {
	return [][]Coordinate {
		{{x: 2, y: 18}, {x: -2, y: 15}},
		{{x: 9, y: 16}, {x: 10, y: 16}},
		{{x: 13, y: 2}, {x: 15, y: 3}},
		{{x: 12, y: 14}, {x: 10, y: 16}},
		{{x: 10, y: 20}, {x: 10, y: 16}},
		{{x: 14, y: 17}, {x: 10, y: 16}},
		{{x: 8, y: 7}, {x: 2, y: 10}},
		{{x: 2, y: 0}, {x: 2, y: 10}},
		{{x: 0, y: 11}, {x: 2, y: 10}},
		{{x: 20, y: 14}, {x: 25, y: 17}},
		{{x: 17, y: 20}, {x: 21, y: 22}},
		{{x: 16, y: 7}, {x: 15, y: 3}},
		{{x: 14, y: 3}, {x: 15, y: 3}},
		{{x: 20, y: 1}, {x: 15, y: 3}},
	}
}
