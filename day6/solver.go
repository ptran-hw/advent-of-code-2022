package day6

import (
	"log"
)

const packetMarkerLength = 4
const messageMarkerLength = 14

type Solver struct {
}

func (s Solver) Solve() {
	//datastreams := getSampleDatastreams()
	datastreams := append([]string{}, readDatastreamFromFile())

	for _, datastream := range datastreams {
		log.Printf("datastream: %s, packet marker starts after character %d\n", datastream,
			findStartMarker(datastream, packetMarkerLength))
		log.Printf("datastream: %s, message marker starts after character %d\n", datastream,
			findStartMarker(datastream, messageMarkerLength))
	}
}

/*
Given []string datastreams, and int windowLength,
Find the index where the windowLength preceding characters are distinct
*/
func findStartMarker(datastream string, windowLength int) int {
	charCount := make(map[string]int, 0)

	head := 0
	tail := 0
	for tail - head < windowLength {
		charToAdd := string(datastream[tail])

		if charCount[charToAdd] > 0 {
			charToRemove := string(datastream[head])
			charCount[charToRemove]--

			head++
			continue
		}

		charCount[charToAdd]++
		tail++

		if tail == len(datastream) {
			log.Panicf("datastream: %s does not have a start marker with length: %d", datastream, windowLength)
		}
	}

	return tail
}
