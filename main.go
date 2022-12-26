package main

import (
	"fmt"
	"os"

	"github.com/ptran-hw/advent-of-code/day1"
)

var solvers = map[string]Solver{
	"1": day1.Solver{},
}

type Solver interface {
	Solve()
}

func main() {
	arguments := os.Args[1:]
	fmt.Printf("Arguments: %s\n", arguments)

	if len(arguments) != 1 {
		panic(fmt.Errorf("incorrect number of arguments used"))
	}

	problemNumber := arguments[0]
	solver := solvers[problemNumber]
	if solver == nil {
		panic(fmt.Errorf("unable to find day%s solver"))
	}

	solver.Solve()
}
