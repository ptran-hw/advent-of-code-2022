package main

import (
	"github.com/ptran-hw/advent-of-code/day1"
	"github.com/ptran-hw/advent-of-code/day10"
	"github.com/ptran-hw/advent-of-code/day11"
	"github.com/ptran-hw/advent-of-code/day12"
	"github.com/ptran-hw/advent-of-code/day13"
	"github.com/ptran-hw/advent-of-code/day14"
	"github.com/ptran-hw/advent-of-code/day15"
	"github.com/ptran-hw/advent-of-code/day16"
	day16_5 "github.com/ptran-hw/advent-of-code/day16.5"
	"github.com/ptran-hw/advent-of-code/day2"
	"github.com/ptran-hw/advent-of-code/day3"
	"github.com/ptran-hw/advent-of-code/day4"
	"github.com/ptran-hw/advent-of-code/day5"
	"github.com/ptran-hw/advent-of-code/day6"
	"github.com/ptran-hw/advent-of-code/day7"
	"github.com/ptran-hw/advent-of-code/day8"
	"github.com/ptran-hw/advent-of-code/day9"
	"log"
	"os"
	"time"
)

// use pointer to Solver for mutable instance
var solvers = map[string]Solver{
	"1":    day1.Solver{},
	"2":    day2.Solver{},
	"3":    day3.Solver{},
	"4":    day4.Solver{},
	"5":    day5.Solver{},
	"6":    day6.Solver{},
	"7":    day7.Solver{},
	"8":    day8.Solver{},
	"9":    day9.Solver{},
	"10":   day10.Solver{},
	"11":   day11.Solver{},
	"12":   day12.Solver{},
	"13":   day13.Solver{},
	"14":   day14.Solver{},
	"15":   day15.Solver{},
	"16":   &day16.Solver{},
	"16.5": day16_5.Solver{},
}

type Solver interface {
	Solve()
}

func main() {
	arguments := os.Args[1:]

	if len(arguments) != 1 {
		log.Panic("incorrect number of arguments used")
	}

	problemNumber := arguments[0]
	solver := solvers[problemNumber]
	if solver == nil {
		log.Panicf("unable to find day %s solver", problemNumber)
	}

	start := time.Now()
	defer func() { log.Printf("time elapsed: %v\n", time.Now().Sub(start)) }()

	solver.Solve()
}
