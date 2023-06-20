package day16

import (
	"fmt"
	"log"
)

type Solver struct {
}

type Room struct {
	label           string
	valveFlowRate   int
	neighbourLabels []string
}

/*
approaches:
- backtrack using roomMap
  - time complexity: c^n exponential
- backtrack using distanceMap
  - time complexity: n! because we have n-1 choices for second room, then n-2 choices for third room, etc
  - room for improvement: the room map is sparse, so there is many rooms with 0 valve pressure, we can skip them
- greedy does not work because it does not consider close proximity for 2nd best choice
  - example:
  - AA, 0 -> BB, CC
  - BB, 10
  - CC, 9 -> DD
  - DD, 9
- using bfs doesnt work because best solution may involve jumping back out and go through another route
  - example:
  - AA, 0 -> BB, CC
  - BB, 10 -> DD, EE
  - CC, 5
  - DD, 0
  - EE, 0
*/
func (s Solver) Solve() {
	const startingRoom = "AA"
	const timeLimit = 30

	//roomMap := getSampleVolcanoRoomMap()
	roomMap := readVolcanoRoomMap()
	distanceMap := buildDistanceMap(roomMap)
	fmt.Sprintf("%v", distanceMap)

	maxPressure := solveMaxPressureReleaseOptimized(startingRoom, timeLimit, roomMap, distanceMap, make(map[string]bool, 0))
	log.Printf("max pressure starting at: %s is %d\n", startingRoom, maxPressure)
}

func buildDistanceMap(roomMap map[string]Room) map[string]int {
	distanceMap := make(map[string]int, 0)
	for _, startRoom := range roomMap {
		putRoomDistances(startRoom, roomMap, distanceMap)
	}

	return distanceMap
}

// uses breadth-first search
func putRoomDistances(startRoom Room, roomMap map[string]Room, distanceMap map[string]int) {
	type QueueEntry struct {
		currRoom Room
		distance int
	}

	queue := make([]QueueEntry, 0)
	buffer := make([]QueueEntry, 0)

	queue = append(queue, QueueEntry{currRoom: startRoom, distance: 0})
	for {
		//log.Printf("queue: %v\n\n", queue)

		for len(queue) > 0 {
			queueEntry := queue[0]
			queue = queue[1:]

			currDistance := queueEntry.distance
			currRoom := queueEntry.currRoom
			distanceMap[getKey(startRoom, currRoom)] = currDistance

			for _, nextRoomLabel := range currRoom.neighbourLabels {
				nextRoom := roomMap[nextRoomLabel]

				if _, exists := distanceMap[getKey(startRoom, nextRoom)]; exists {
					continue
				}

				buffer = append(buffer, QueueEntry{currRoom: nextRoom, distance: currDistance + 1})
			}
		}

		if len(buffer) == 0 {
			break
		}

		queue = append(queue, buffer...) // nice syntax sugar
		buffer = make([]QueueEntry, 0)
	}
}

func getKey(roomA, roomB Room) string {
	return fmt.Sprintf("%s,%s", roomA.label, roomB.label)
}

// TODO: refactor to remove the constant parameters into struct
func solveMaxPressureReleaseOptimized(startingRoomLabel string, timeLimit int, roomMap map[string]Room,
	distanceMap map[string]int, visitedRooms map[string]bool) int {
	if timeLimit <= 0 {
		return 0
	}

	visitedRooms[startingRoomLabel] = true
	currRoom := roomMap[startingRoomLabel]

	maxPressure := 0
	for nextRoomLabel, _ := range roomMap {
		if visitedRooms[nextRoomLabel] {
			continue
		}

		visitedRooms[nextRoomLabel] = true

		nextRoom := roomMap[nextRoomLabel]
		// this change increased performance from 170ms to 3ms for sample input (size 10)
		if nextRoom.valveFlowRate == 0 {
			continue
		}

		nextTimeLimit := timeLimit - distanceMap[getKey(currRoom, nextRoom)] - 1
		// used to have check for negative next time limit, but we are always caught in base case and the resulting pressure would be negative
		// it's a bit implicit so for code readability, uncertain if it's better

		nextRoomPressureContribution := nextRoom.valveFlowRate * nextTimeLimit
		maxPressureFromRecursion := solveMaxPressureReleaseOptimized(nextRoomLabel, nextTimeLimit, roomMap, distanceMap, visitedRooms)
		visitedRooms[nextRoomLabel] = false

		if maxPressure < nextRoomPressureContribution + maxPressureFromRecursion {
			maxPressure = nextRoomPressureContribution + maxPressureFromRecursion
		}
	}

	return maxPressure
}

// Part 1
// approach uses backtrack to explore all the different paths, considering opening the current room valve and not opening
func solveMaxPressureRelease(startingRoomLabel string, timeLimit int, roomMap map[string]Room) {
	maxPressure := solveMaxPressureReleaseHelper(startingRoomLabel, timeLimit, roomMap, make(map[string]bool, 0))
	log.Printf("For timelimit %d, the max pressure release is: %d\n", timeLimit, maxPressure)
}


func solveMaxPressureReleaseHelper(currentRoomLabel string, timeLimit int, roomMap map[string]Room, valveOpenedRooms map[string]bool) int {
	if timeLimit <= 0 {
		return 0
	}

	maxPressureReleased := 0
	for _, nextRoomLabel := range roomMap[currentRoomLabel].neighbourLabels {
		maxPressureReleasedWithoutCurrentValveOpened := solveMaxPressureReleaseHelper(nextRoomLabel, timeLimit - 1, roomMap, valveOpenedRooms)
		if maxPressureReleased < maxPressureReleasedWithoutCurrentValveOpened {
			maxPressureReleased = maxPressureReleasedWithoutCurrentValveOpened
		}

		currentRoom := roomMap[currentRoomLabel]
		pressureReleasedOverTime := currentRoom.valveFlowRate * (timeLimit - 1)

		valveOpenedRooms[currentRoomLabel] = true
		maxPressureReleasedWithCurrentValveOpened := solveMaxPressureReleaseHelper(nextRoomLabel, timeLimit - 2, roomMap, valveOpenedRooms) + pressureReleasedOverTime
		if maxPressureReleased < maxPressureReleasedWithCurrentValveOpened {
			maxPressureReleased = maxPressureReleasedWithCurrentValveOpened
		}
		valveOpenedRooms[currentRoomLabel] = false
	}

	return maxPressureReleased
}