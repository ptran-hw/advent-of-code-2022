package day13

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// Notes:
// - found alot of bugs created from regex matching: i) ,? ii) ^
// - tried to use regex to parse directly, but could not do it for recursive data structures
// - we can cast string(currLine[index]) or return string directly currLine[index:index]
// - also tested in main.go to verify behaviour

const signalFile string = "./day13/signals.txt"

var dataListRegex = regexp.MustCompile("\\[(.*)\\]")
var dataPointRegex = regexp.MustCompile("^(\\d+),?(.*)")

func readSignalsFromFile() []Signal {
	file, err := os.Open(signalFile)
	if err != nil {
		panic(err)
	}

	signals := make([]Signal, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		signal := Signal{values: parseSignalValues(line)}
		signals = append(signals, signal)
	}

	return signals
}

// example lines:
// [1,2]
// [[1]]
// [1,[2]]
func parseSignalValues(line string) []SignalValue {
	if !dataListRegex.MatchString(line) {
		return nil
	}

	results := make([]SignalValue, 0)
	currLine := dataListRegex.FindStringSubmatch(line)[1]
	for len(currLine) > 0 {
		switch {
		case dataPointRegex.MatchString(currLine):
			matches := dataPointRegex.FindStringSubmatch(currLine)

			intValue, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}

			results = append(results, SignalValue{
				dataPoint: intValue,
			})

			currLine = matches[2]
		default: // data list scenario
			buffer := "["
			bracketCounter := 1
			index := 1

			for bracketCounter > 0 {
				buffer = buffer + string(currLine[index])

				if string(currLine[index]) == "[" {
					bracketCounter++
				} else if string(currLine[index]) == "]" {
					bracketCounter--
				}

				index++
			}

			// recurse and get result of smaller data list
			subResult := parseSignalValues(buffer)
			results = append(results, SignalValue{
				dataList: subResult,
			})
			currLine = currLine[index:]
			if len(currLine) > 0 && string(currLine[0]) == "," {
				currLine = currLine[1:]
			}
		}
	}

	return results
}

func getSampleSignals() []Signal {
	return []Signal{
		{
			values: []SignalValue{
				{dataPoint: 1},
				{dataPoint: 1},
				{dataPoint: 3},
				{dataPoint: 1},
				{dataPoint: 1},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 1},
				{dataPoint: 1},
				{dataPoint: 5},
				{dataPoint: 1},
				{dataPoint: 1},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{dataPoint: 1},
					},
				},
				{
					dataList: []SignalValue{
						{dataPoint: 2},
						{dataPoint: 3},
						{dataPoint: 4},
					},
				},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{dataPoint: 1},
					},
				},
				{dataPoint: 4},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 9},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{dataPoint: 8},
						{dataPoint: 7},
						{dataPoint: 6},
					},
				},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{dataPoint: 4},
						{dataPoint: 4},
					},
				},
				{dataPoint: 4},
				{dataPoint: 4},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{dataPoint: 4},
						{dataPoint: 4},
					},
				},
				{dataPoint: 4},
				{dataPoint: 4},
				{dataPoint: 4},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 7},
				{dataPoint: 7},
				{dataPoint: 7},
				{dataPoint: 7},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 7},
				{dataPoint: 7},
				{dataPoint: 7},
			},
		},
		{
			values: []SignalValue{},
		},
		{
			values: []SignalValue{
				{dataPoint: 3},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{
							dataList: []SignalValue{
								{
									dataList: []SignalValue{},
								},
							},
						},
					},
				},
			},
		},
		{
			values: []SignalValue{
				{
					dataList: []SignalValue{
						{
							dataList: []SignalValue{},
						},
					},
				},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 1},
				{
					dataList: []SignalValue{
						{dataPoint: 2},
						{
							dataList: []SignalValue{
								{dataPoint: 3},
								{
									dataList: []SignalValue{
										{dataPoint: 4},
										{
											dataList: []SignalValue{
												{dataPoint: 5},
												{dataPoint: 6},
												{dataPoint: 7},
											},
										},
									},
								},
							},
						},
					},
				},
				{dataPoint: 8},
				{dataPoint: 9},
			},
		},
		{
			values: []SignalValue{
				{dataPoint: 1},
				{
					dataList: []SignalValue{
						{dataPoint: 2},
						{
							dataList: []SignalValue{
								{dataPoint: 3},
								{
									dataList: []SignalValue{
										{dataPoint: 4},
										{
											dataList: []SignalValue{
												{dataPoint: 5},
												{dataPoint: 6},
												{dataPoint: 0},
											},
										},
									},
								},
							},
						},
					},
				},
				{dataPoint: 8},
				{dataPoint: 9},
			},
		},
	}
}

// brainstorming:
// - start with "[", create a parent signalValue with data list (maybe we need to recurse)
// - see numbers, then recurse which will create signal value data point
// - see "]", then parent signal value is done
// re-iterate idea:
// - use regex to extract the values within []
// - check if matches: num, () then we can break it down
// - check if matches: [()], but we may face issues with [],[]
// maybe we can consider stack and spitting out sub array when see "]"
// - stack eats up every token ("[", "numbers")
// - when we iterate to a "]", then we pop from stack until we see "["
// - recreate string and call helper function
// - repeat, so we are only dealing with 1-d arrays at any moment
// - problem is when we have nested relationship we cannot handle iteratively easily
// another approach:
// - ignore first open and closed bracket, can use regex pattern: ^[(inside)]$
// - start with number
//	 - match regexp pattern: (number), (rest)
// - start with open bracket
// - find matching closed bracket
// - recurse on substring