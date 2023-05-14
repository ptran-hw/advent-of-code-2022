package main

import (
	"fmt"
	"github.com/ptran-hw/advent-of-code/day1"
	"github.com/ptran-hw/advent-of-code/day10"
	"github.com/ptran-hw/advent-of-code/day11"
	"github.com/ptran-hw/advent-of-code/day2"
	"github.com/ptran-hw/advent-of-code/day3"
	"github.com/ptran-hw/advent-of-code/day4"
	"github.com/ptran-hw/advent-of-code/day5"
	"github.com/ptran-hw/advent-of-code/day6"
	"github.com/ptran-hw/advent-of-code/day7"
	"github.com/ptran-hw/advent-of-code/day8"
	"github.com/ptran-hw/advent-of-code/day9"
	"os"
)

// use pointer to Solver for mutable instance
var solvers = map[string]Solver{
	"1": day1.Solver{},
	"2": day2.Solver{},
	"3": day3.Solver{},
	"4": day4.Solver{},
	"5": day5.Solver{},
	"6": day6.Solver{},
	"7": day7.Solver{},
	"8": day8.Solver{},
	"9": day9.Solver{},
	"10": &day10.Solver{},
	"11": &day11.Solver{},
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
