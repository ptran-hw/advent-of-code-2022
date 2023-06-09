package day15

import (
	"log"
	"math"
	"sort"
)

type Coordinate struct {
	x int
	y int
}

type Range struct {
	start int
	end int
}

type Solver struct {
}

const beaconCharacter = "B"
const nonBeaconCharacter = "N"

func (s Solver) Solve() {
	coordinatePairs, rowIndex, distressBeaconLowerBound, distressBeaconUpperBound := getSampleSensorBeaconCoordinates(), 10, 0, 20
	//coordinatePairs, rowIndex, distressBeaconLowerBound, distressBeaconUpperBound := readSensorBeaconCoordinates(), 2_000_000, 0, 4_000_000

	positionCount, beaconCount, _ := solveBeaconExclusionUsingRangeApproach(coordinatePairs, rowIndex, math.MinInt, math.MaxInt)
	log.Printf("At y=%d, there is %d non-beacon positions (%d detected positions and %d beacons)\n", rowIndex, positionCount - beaconCount, positionCount, beaconCount)

	solveMissingDistressBeacon(coordinatePairs, distressBeaconLowerBound, distressBeaconUpperBound)
}

// Part 1
// approach involves iterating on x-axis and adding to map[int]bool. At the end we loop over the map to count the result
// for sample data input, roughly same execution time 65 microseconds
// for large data input, roughly same execution time 1 second
func solveBeaconExclusion(coordinatePairs [][]Coordinate, rowIndex int) {
	beaconMap := getBeaconMap(coordinatePairs, rowIndex)

	log.Printf("There is %d non-beacon positions for y=%d\n", countNonBeaconPositions(beaconMap), rowIndex)
}

func getBeaconMap(coordinatePairs [][]Coordinate, rowIndex int) map[int]string {
	visitedMap := make(map[int]string, 0)

	for _, pair := range coordinatePairs {
		sensor := pair[0]
		beacon := pair[1]

		sensorRange := calculateManhattanDistance(sensor, beacon)

		if beacon.y == rowIndex {
			visitedMap[beacon.x] = beaconCharacter
		}

		xOffset := 0
		for calculateManhattanDistance(sensor, Coordinate{x: sensor.x + xOffset, y: rowIndex}) <= sensorRange {
			if _, exists := visitedMap[sensor.x+xOffset]; !exists {
				visitedMap[sensor.x+xOffset] = nonBeaconCharacter
			}

			if _, exists := visitedMap[sensor.x-xOffset]; !exists {
				visitedMap[sensor.x-xOffset] = nonBeaconCharacter
			}

			xOffset++
		}
	}
	return visitedMap
}

func calculateManhattanDistance(coordinateA, coordinateB Coordinate) int {
	floatResult := math.Abs(float64(coordinateA.x) - float64(coordinateB.x)) +
		math.Abs(float64(coordinateA.y) - float64(coordinateB.y))
	return int(floatResult)
}

func countNonBeaconPositions(visitedMap map[int]string) int {
	count := 0
	for _, value := range visitedMap {
		if value == nonBeaconCharacter {
			count++
		}
	}

	return count
}

// Part 1 optimized
// for sample data input, roughly same execution time 44 microseconds
// for large data input, an improvement from 1s to 200 microseconds
func solveBeaconExclusionUsingRangeApproach(coordinatePairs [][]Coordinate, rowIndex, lowerBound, upperBound int) (int, int, []Range) {
	ranges := getBeaconRanges(coordinatePairs, rowIndex, lowerBound, upperBound)
	combinedRanges := combineRanges(ranges)
	positionCount := countTotalLength(combinedRanges)
	beaconCount := countBeacons(coordinatePairs, rowIndex)

	return positionCount, beaconCount, combinedRanges
}

// provides windows of coverage using manhattan distance
func getBeaconRanges(coordinatePairs [][]Coordinate, rowIndex, lowerBound, upperBound int) []Range {
	coveredRanges := make([]Range, 0)

	for _, pair := range coordinatePairs {
		sensor := pair[0]
		beacon := pair[1]

		// sensor to beacon distance
		availableManhattanDistance := calculateManhattanDistance(sensor, beacon)

		// sensor to potential beacons distance
		var yDiff int
		if sensor.y >= rowIndex {
			yDiff = sensor.y - rowIndex
		} else {
			yDiff = rowIndex - sensor.y
		}

		remainingManhattanDistance := availableManhattanDistance - yDiff
		if remainingManhattanDistance <= 0 {
			continue
		}

		// for part 2
		startValue := sensor.x - remainingManhattanDistance
		if startValue < lowerBound {
			startValue = lowerBound
		}

		endValue := sensor.x + remainingManhattanDistance
		if endValue > upperBound {
			endValue = upperBound
		}

		coveredRanges = append(coveredRanges, Range{start: startValue, end: endValue})
	}

	return coveredRanges
}

// we want to combine:
// - [0, 1], [1, 2] -> [0, 2]
// - [1, 3], [2, 5] -> [1, 5]
func combineRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		rangeA := ranges[i]
		rangeB := ranges[j]

		if rangeA.start != rangeB.start {
			return rangeA.start < rangeB.start
		}

		return rangeA.end < rangeB.end
	})

	combinedRange := ranges[0]
	result := make([]Range, 0)
	for index, currRange := range ranges {
		if !canCombine(combinedRange, currRange) {
			result = append(result, combinedRange)
			combinedRange = currRange
		}

		if combinedRange.start > currRange.start {
			combinedRange.start = currRange.start
		}
		if combinedRange.end < currRange.end {
			combinedRange.end = currRange.end
		}

		if index == len(ranges) - 1 {
			result = append(result, combinedRange)
		}
	}

	return result
}

func countTotalLength(ranges []Range) int {
	total := 0
	for _, currRange := range ranges {
		total += currRange.end - currRange.start + 1
	}

	return total
}

// assume rangeA and rangeB is sorted order
func canCombine(rangeA, rangeB Range) bool {
	return isOverlapping(rangeA, rangeB) || rangeA.end + 1 == rangeB.start
}

func isOverlapping(rangeA, rangeB Range) bool {
	isNonOverlapping := rangeA.end < rangeB.start || rangeA.start > rangeB.end
	return !isNonOverlapping
}

func countBeacons(coordinatePairs [][]Coordinate, rowIndex int) int {
	visited := make(map[int]bool, 0)

	for _, pair := range coordinatePairs {
		beacon := pair[1]

		if beacon.y == rowIndex {
			visited[beacon.x] = true
		}
	}

	return len(visited)
}

// Part 2
func solveMissingDistressBeacon(coordinatePairs [][]Coordinate, lowerBound, upperBound int) {
	gridWidth := upperBound - lowerBound + 1
	for rowIndex := lowerBound; rowIndex <= upperBound; rowIndex++ {
		detectedPositionCount, _, detectedRanges := solveBeaconExclusionUsingRangeApproach(coordinatePairs, rowIndex, lowerBound, upperBound)
		if detectedPositionCount < gridWidth {
			xPosition := findMissingPosition(detectedRanges, lowerBound, upperBound)
			log.Printf("tuning frequency for (%d, %d) is: %d\n", xPosition, rowIndex, calculateTuningFrequency(xPosition, rowIndex))
		}
	}
}

func findMissingPosition(sortedRanges []Range, lowerBound, upperBound int) int {
	if len(sortedRanges) > 2 {
		panic("there should be either 1 or 2 ranges because there is at most one solution")
	}

	if sortedRanges[0].start != lowerBound {
		return lowerBound
	}
	if sortedRanges[len(sortedRanges) - 1].end != upperBound {
		return upperBound
	}

	return sortedRanges[0].end + 1
}

func calculateTuningFrequency(x, y int) int {
	return x * 4_000_000 + y
}

// Attempt 1: looping over x = 0 to 4_000_000
// we would start at rowIndex and iterate with xOffset = 1, 2, ...
// this would run 4_000_000 times for grid size 4_000_000