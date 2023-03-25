package main

import (
	"fmt"
	"github.com/ptran-hw/advent-of-code/day2"
	"github.com/ptran-hw/advent-of-code/day3"
	"github.com/ptran-hw/advent-of-code/day4"
	"os"

	"github.com/ptran-hw/advent-of-code/day1"
)

var solvers = map[string]Solver{
	"1": day1.Solver{},
	"2": day2.Solver{},
	"3": day3.Solver{},
	"4": day4.Solver{},
}

type Solver interface {
	Solve()
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		panic(fmt.Errorf("incorrect number of arguments used"))
	}

	problemNumber := arguments[0]
	solver := solvers[problemNumber]
	if solver == nil {
		panic(fmt.Errorf("unable to find day %s solver", problemNumber))
	}

	solver.Solve()
}
