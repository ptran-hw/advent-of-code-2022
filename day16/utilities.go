package day16

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const tunnelDataFile = "./day16/tunnelData.txt"
var valveTunnelRegexp = regexp.MustCompile("Valve (.+) has flow rate=(\\d+); tunnels? leads? to valves? (.+)")

func readVolcanoRoomMap() map[string]Room {
	file, err := os.Open(tunnelDataFile)
	if err != nil {
		panic(err)
	}

	roomMap := make(map[string]Room, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if !valveTunnelRegexp.MatchString(line) {
			log.Fatalln("unable to parse line:", line)
		}

		tokens := valveTunnelRegexp.FindStringSubmatch(line)
		roomLabel := tokens[1]
		flowRate := tokens[2]
		neighbourRooms := tokens[3]

		flowRateValue, err := strconv.Atoi(flowRate)
		if err != nil {
			panic(err)
		}

		neighbourRoomsValue := strings.Split(neighbourRooms, ", ")

		roomMap[roomLabel] = Room{
			label: roomLabel,
			valveFlowRate: flowRateValue,
			neighbourLabels: neighbourRoomsValue,
		}
	}

	return roomMap
}

func getSampleVolcanoRoomMap() map[string]Room {
	labelRoomMap := map[string]Room {
		"AA": {
			label: "AA",
			valveFlowRate: 0,
			neighbourLabels: []string{"DD", "II", "BB"},
		},
		"BB": {
			label: "BB",
			valveFlowRate: 13,
			neighbourLabels: []string{"CC", "AA"},
		},
		"CC": {
			label: "CC",
			valveFlowRate: 2,
			neighbourLabels: []string{"DD", "BB"},
		},
		"DD": {
			label: "DD",
			valveFlowRate: 20,
			neighbourLabels: []string{"CC", "AA", "EE"},
		},
		"EE": {
			label: "EE",
			valveFlowRate: 3,
			neighbourLabels: []string{"FF", "DD"},
		},
		"FF": {
			label: "FF",
			valveFlowRate: 0,
			neighbourLabels: []string{"EE", "GG"},
		},
		"GG": {
			label: "GG",
			valveFlowRate: 0,
			neighbourLabels: []string{"FF", "HH"},
		},
		"HH": {
			label: "HH",
			valveFlowRate: 22,
			neighbourLabels: []string{"GG"},
		},
		"II": {
			label: "II",
			valveFlowRate: 0,
			neighbourLabels: []string{"AA", "JJ"},
		},
		"JJ": {
			label: "JJ",
			valveFlowRate: 21,
			neighbourLabels: []string{"II"},
		},
	}

	return labelRoomMap
}
