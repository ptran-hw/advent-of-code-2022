package day13

import (
	"fmt"
	"reflect"
	"sort"
)

const indexOffset int = 1
const isInOrder int = 1
const isEqual int = 0
const isNotInOrder int = -1

var decoderKeyA = getDecoderKey(2)
var decoderKeyB = getDecoderKey(6)

type Signal struct {
	values []SignalValue
}

type SignalValue struct {
	dataPoint int
	dataList []SignalValue
}

func (sv SignalValue) isDataPoint() bool {
	return !sv.isDataList()
}

func (sv SignalValue) isDataList() bool {
	return sv.dataList != nil
}

type Solver struct {
}

func (s Solver) Solve() {
	//signals := getSampleSignals()
	signals := readSignalsFromFile()

	solveInOrderScore(signals)

	signals = append(signals, decoderKeyA)
	signals = append(signals, decoderKeyB)
	sortSignals(signals)
	solveDecoderScore(signals)
}

func solveInOrderScore(signals []Signal) {
	score := 0
	for pairIndex := 0; pairIndex < len(signals) / 2; pairIndex++ {
		signalA, signalB := getSignalPair(pairIndex, signals)

		orderState := compare(signalA.values, signalB.values)
		if orderState == isEqual || orderState == isInOrder {
			score += pairIndex + indexOffset
		}
	}

	fmt.Printf("The in-order score of signal pairs is: %d\n", score)
}

func getSignalPair(pairIndex int, signals []Signal) (Signal, Signal) {
	return signals[pairIndex * 2], signals[(pairIndex * 2) + 1]
}

func compare(signalValuesA, signalValuesB []SignalValue) int {
	for index := 0; index < len(signalValuesA) && index < len(signalValuesB); index++ {
		left := signalValuesA[index]
		right := signalValuesB[index]

		switch {
		case left.isDataPoint() && right.isDataPoint():
			if left.dataPoint < right.dataPoint {
				return isInOrder
			} else if left.dataPoint > right.dataPoint {
				return isNotInOrder
			}

			continue
		case left.isDataPoint() && !right.isDataPoint():
			leftDataList := []SignalValue{
				{dataPoint: left.dataPoint},
			}

			result := compare(leftDataList, right.dataList)
			if result != isEqual {
				return result
			}

			continue
		case !left.isDataPoint() && right.isDataPoint():
			rightDataList := []SignalValue{
				{dataPoint: right.dataPoint},
			}

			result := compare(left.dataList, rightDataList)
			if result != isEqual {
				return result
			}

			continue
		case !left.isDataPoint() && !right.isDataPoint():
			result := compare(left.dataList, right.dataList)
			if result != isEqual {
				return result
			}

			continue
		default:
			panic("Unable to compare signal values")
		}
	}

	if len(signalValuesA) < len(signalValuesB) {
		return isInOrder
	} else if len(signalValuesA) > len(signalValuesB) {
		return isNotInOrder
	}

	return isEqual
}

func getDecoderKey(value int) Signal {
	return Signal{
		values: []SignalValue{
			{
				dataList: []SignalValue{
					{dataPoint: value},
				},
			},
		},
	}
}

func sortSignals(signals []Signal) {
	sort.Slice(signals, func(i, j int) bool { return compare(signals[i].values, signals[j].values) >= 0 })
}

func solveDecoderScore(signals []Signal) {
	indexA, indexB := 0, 0
	for index := 0; index < len(signals); index++ {
		switch {
		case reflect.DeepEqual(signals[index], decoderKeyA):
			indexA = index
		case reflect.DeepEqual(signals[index], decoderKeyB):
			indexB = index
		}
	}

	fmt.Printf("Decoders found at index: %d, %d and decoder score is %d\n", indexA, indexB, (indexA + 1) * (indexB + 1))
}