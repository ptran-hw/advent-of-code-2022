package day6

import (
	"bufio"
	"log"
	"os"
)

const datastreamInputFile = "./day6/datastreamInput.txt"

func readDatastreamFromFile() string {
	file, err := os.Open(datastreamInputFile)
	if err != nil {
		log.Panicf("unable to read datastream file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Panic("invalid datastream file, must contain a single line")
	}

	return scanner.Text()
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
