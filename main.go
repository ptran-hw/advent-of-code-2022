\package main

import (
	"fmt"
	"github.com/ptran-hw/advent-of-code/day2"
	"os"

	"github.com/ptran-hw/advent-of-code/day1"
)

var solvers = map[string]Solver{
	"1":   day1.Solver{},
	"2":   day2.Solver{},
	"2.1": day2.Solver{},
	"2.2": day2.SequelSolver{},
}

type Solver interface {
	Solve()
}

func main() {
	arguments := os.Args[1:]
	arguments = []string{"2.2"}

	if len(arguments) != 1 {
		panic(fmt.Errorf("incorrect number of arguments used"))
	}

	problemNumber := arguments[0]
	solver := solvers[problemNumber]
	if solver == nil {
		panic(fmt.Errorf("unable to find day%s solver", problemNumber))
	}

	solver.Solve()
}
