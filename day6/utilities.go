package day6

import (
	"bufio"
	"os"
)

const datastreamInputFile = "./day6/datastreamInput.txt"

func readDatastreamFromFile() []string {
	file, err := os.Open(datastreamInputFile)
	if err != nil {
		panic(err)
	}

	datastreams := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		datastreams = append(datastreams, line)
	}

	return datastreams
}

func getSampleDatastreams() []string {
	return []string{
		"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		"bvwbjplbgvbhsrlpgdmjqwftvncz",
		"nppdvjthqldpwncqszvftbrmjlhg",
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
	}
}
