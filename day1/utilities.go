package day1

import (
	"bufio"
	"math"
	"os"
	"strconv"

	"github.com/agrison/go-commons-lang/stringUtils"
)

func ReadElfCaloriesFromInputFile() [][]int {
	file, err := os.Open(InputFile)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var result [][]int
	var buffer []int
	for scanner.Scan() {
		if stringUtils.IsNotEmpty(scanner.Text()) {
			value, err := strconv.Atoi(scanner.Text())
			if err != nil {
				panic(err)
			}

			buffer = append(buffer, value)
		} else {
			result = append(result, buffer)
			buffer = make([]int, 0)
		}
	}

	return result
}

func GetSampleElfCalories() [][]int {
	return [][]int{
		{1000, 2000, 3000},
		{4000},
		{5000, 6000},
		{7000, 8000, 9000},
		{10000},
	}
}

func GetMaxValue(values []int) int {
	return GetMaxValueWithUpperbound(values, math.MaxInt)
}

func GetMaxValueWithUpperbound(values []int, upperbound int) int {
	maxValue := 0
	for _, value := range values {
		if value <= upperbound {
			maxValue = Max(maxValue, value)
		}
	}

	return maxValue
}

func Max(valA, valB int) int {
	return int(math.Max(float64(valA), float64(valB)))
}
