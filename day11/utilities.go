package day11

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const monkeyDataFile = "./day11/monkeyData.txt"

var monkeyHeader = regexp.MustCompile("Monkey")
var worryLevelsPattern = regexp.MustCompile("Starting items: (.+)$")
var operationFuncPattern = regexp.MustCompile("Operation: new = old ([\\+\\-\\*\\/\\^]) (.+)$")
var testFuncDivisiblePattern = regexp.MustCompile("Test: divisible by (.+)$")
var testFuncTruePattern = regexp.MustCompile("If true: throw to monkey (.+)$")
var testFuncFalsePattern = regexp.MustCompile("If false: throw to monkey (.+)$")

func scanNextLine(scanner *bufio.Scanner) {
	if !scanner.Scan() {
		panic("invalid monkey data")
	}
}

func readMonkeysFromFile() []*Monkey {
	file, err := os.Open(monkeyDataFile)
	if err != nil {
		panic(err)
	}

	monkeys := make([]*Monkey, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !monkeyHeader.MatchString(line) {
			// ignore lines until we find a valid header
			continue
		}

		worryLevels := getWorryLevels(scanner)
		operationFunc := getOperationFunc(scanner)
		redirectTestFunc, divisibleValue := getRedirectTestFunc(scanner)
		newMonkey := &Monkey{
			worryLevels:      worryLevels,
			operationFunc:    operationFunc,
			redirectTestFunc: redirectTestFunc,
			redirectDivisibleValue: divisibleValue,
		}

		monkeys = append(monkeys, newMonkey)
	}

	return monkeys
}

func getWorryLevels(scanner *bufio.Scanner) []int {
	scanNextLine(scanner)
	line := scanner.Text()

	if !worryLevelsPattern.MatchString(line) {
		panic(fmt.Sprintf("Expected starting items line, but instead got: %s", line))
	}

	match := worryLevelsPattern.FindStringSubmatch(line)[1]
	stressLevels := make([]int, 0)
	for _, token := range strings.Split(match, ", ") {
		value, err := strconv.Atoi(token)
		if err != nil {
			panic(err)
		}

		stressLevels = append(stressLevels, value)
	}

	return stressLevels
}

func getOperationFunc(scanner *bufio.Scanner) func(int) int {
	scanNextLine(scanner)
	line := scanner.Text()

	if !operationFuncPattern.MatchString(line) {
		panic(fmt.Sprintf("Expected operation function line, but instead got: %s", line))
	}

	operator := operationFuncPattern.FindStringSubmatch(line)[1]
	argument := operationFuncPattern.FindStringSubmatch(line)[2]
	if argument == "old" {
		switch operator {
		case "+":
			operator = "*"
			argument = "2"
		case "-":
			operator = "*"
			argument = "0"
		case "*":
			operator = "^"
			argument = "2"
		case "/":
			operator = "^"
			argument = "0"
		}
	}

	argumentValue, err := strconv.Atoi(argument)
	if err != nil {
		panic(err)
	}

	return getOperatorFuncClosure(operator, argumentValue)
}

func getOperatorFuncClosure(operator string, value int) func(int) int {
	funcs := map[string]func(int) int {
		"+": func(x int) int {
			return x + value
		},
		"-": func(x int) int {
			return x - value
		},
		"*": func(x int) int {
			return x * value
		},
		"/": func(x int) int {
			return x / value
		},
		"^": func(x int) int {
			return int(math.Pow(float64(x), float64(value)))
		},
	}

	return funcs[operator]
}

func getRedirectTestFunc(scanner *bufio.Scanner) (func(int) int, int) {
	divisibleValue := getRedirectTestDivisibleValue(scanner)
	truePathMonkey := getRedirectTestTruePathMonkey(scanner)
	falsePathMonkey := getRedirectTestFalsePathMonkey(scanner)

	return func(x int) int {
		if x % divisibleValue == 0 {
			return truePathMonkey
		} else {
			return falsePathMonkey
		}
	}, divisibleValue
}

func getRedirectTestDivisibleValue(scanner *bufio.Scanner) int {
	scanNextLine(scanner)
	line := scanner.Text()

	if !testFuncDivisiblePattern.MatchString(line) {
		panic(fmt.Sprintf("Expected Test function divisible line, but instead got: %s", line))
	}

	divisible := testFuncDivisiblePattern.FindStringSubmatch(line)[1]
	divisibleValue, err := strconv.Atoi(divisible)
	if err != nil {
		panic(err)
	}

	return divisibleValue
}

func getRedirectTestTruePathMonkey(scanner *bufio.Scanner) int {
	scanNextLine(scanner)
	line := scanner.Text()

	if !testFuncTruePattern.MatchString(line) {
		panic(fmt.Sprintf("Expected Test function true line, but instead got: %s", line))
	}

	monkey := testFuncTruePattern.FindStringSubmatch(line)[1]
	monkeyValue, err := strconv.Atoi(monkey)
	if err != nil {
		panic(err)
	}

	return monkeyValue
}

func getRedirectTestFalsePathMonkey(scanner *bufio.Scanner) int {
	scanNextLine(scanner)
	line := scanner.Text()

	if !testFuncFalsePattern.MatchString(line) {
		panic(fmt.Sprintf("Expected Test function false line, but instead got: %s", line))
	}

	monkey := testFuncFalsePattern.FindStringSubmatch(line)[1]
	monkeyValue, err := strconv.Atoi(monkey)
	if err != nil {
		panic(err)
	}

	return monkeyValue
}

func getSampleMonkeys() []*Monkey {
	return []*Monkey {
		{
			worryLevels: []int{79, 98},
			operationFunc: func(old int) int {
				return old * 19
			},
			redirectTestFunc: func(level int) int {
				if level % 23 == 0 {
					return 2
				} else {
					return 3
				}
			},
			redirectDivisibleValue: 23,
		},
		{
			worryLevels: []int{54, 65, 75, 74},
			operationFunc: func(old int) int {
				return old + 6
			},
			redirectTestFunc: func(level int) int {
				if level % 19 == 0 {
					return 2
				} else {
					return 0
				}
			},
			redirectDivisibleValue: 19,
		},
		{
			worryLevels: []int{79, 60, 97},
			operationFunc: func(old int) int {
				return old * old
			},
			redirectTestFunc: func(level int) int {
				if level % 13 == 0 {
					return 1
				} else {
					return 3
				}
			},
			redirectDivisibleValue: 13,
		},
		{
			worryLevels: []int{74},
			operationFunc: func(old int) int {
				return old + 3
			},
			redirectTestFunc: func(level int) int {
				if level % 17 == 0 {
					return 0
				} else {
					return 1
				}
			},
			redirectDivisibleValue: 17,
		},
	}
}
