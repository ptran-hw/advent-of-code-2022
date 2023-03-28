package day6

import "fmt"

const packetMarkerLength = 4
const messageMarkerLength = 14

type Solver struct {
}

func (s Solver) Solve() {
	findPacketStartMarkers()
	findMessageStartMarkers()
}

func findPacketStartMarkers() {
	datastreams := readDatastreamFromFile()
	//datastreams := getSampleDatastreams()

	markers := make([]int, 0)
	for _, datastream := range datastreams {
		markers = append(markers, findStartMarker(datastream, packetMarkerLength))
	}

	for index := 0; index < len(datastreams); index++ {
		fmt.Printf("datastream: %s, message marker starts after character %d\n", datastreams[index], markers[index])
	}
}

func findMessageStartMarkers() {
	datastreams := readDatastreamFromFile()
	//datastreams := getSampleDatastreams()

	markers := make([]int, 0)
	for _, datastream := range datastreams {
		markers = append(markers, findStartMarker(datastream, messageMarkerLength))
	}

	for index := 0; index < len(datastreams); index++ {
		fmt.Printf("datastream: %s, message marker starts after character %d\n", datastreams[index], markers[index])
	}
}

func findStartMarker(datastream string, markerLength int) int {
	charCount := make(map[string]int, 0)

	head := 0
	tail := 0
	for tail-head < messageMarkerLength {
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
			panic(fmt.Sprintf("datastream: %s does not have a start marker with length: %d", datastream, markerLength))
		}
	}

	return tail
}
