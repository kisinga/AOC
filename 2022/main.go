//package main is the entry point for all our solutions
//All the code in this file is honestly unecessary, as I could have easily run the solutions independently, but it comes with some advantages too
//1. Standardization of the function calls
//2. I don't have to cd into a different folder every time I want to run a different solution
//3. Just learning really. It's an opportunity to introduce some structure into the code which just demonstrates some prowess

package main

import (
	"flag"
	"fmt"
	. "github.com/kisinga/AOC/2022/day01"
	"os"
	"reflect"
	"strconv"
)

// This method must be implemented by all our solutions for them to be callable from a central place
type runner interface {
	Run()
}

// This struct contains all our solutions functions in the form of the Runner method
type solutions struct {
	Day01 runner
	Day02 runner
	Day03 runner
	Day04 runner
	Day05 runner
	Day06 runner
	Day07 runner
	Day08 runner
	Day09 runner
	Day10 runner
	Day11 runner
	Day12 runner
	Day13 runner
	Day14 runner
	Day15 runner
	Day16 runner
	Day17 runner
	Day18 runner
	Day19 runner
	Day20 runner
	Day21 runner
	Day22 runner
	Day23 runner
	Day24 runner
	Day25 runner
	Day26 runner
	Day27 runner
	Day28 runner
	Day29 runner
	Day30 runner
}

func NewSolutions() solutions {
	return solutions{
		Day01: NewDay01(),
		Day02: nil,
		Day03: nil,
		Day04: nil,
		Day05: nil,
		Day06: nil,
		Day07: nil,
		Day08: nil,
		Day09: nil,
		Day10: nil,
		Day11: nil,
		Day12: nil,
		Day13: nil,
		Day14: nil,
		Day15: nil,
		Day16: nil,
		Day17: nil,
		Day18: nil,
		Day19: nil,
		Day20: nil,
		Day21: nil,
		Day22: nil,
		Day23: nil,
		Day24: nil,
		Day25: nil,
		Day26: nil,
		Day27: nil,
		Day28: nil,
		Day29: nil,
		Day30: nil,
	}

}

func main() {
	exitCode := 0
	defer func() { os.Exit(exitCode) }()
	solution := NewSolutions()

	//get our cli args
	flag.Parse()
	day := flag.Args()[0]
	//make sure that the day is within the calendar range
	dayInt, err := strconv.Atoi(day)
	if dayInt > 30 || err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", "Invalid day provided. Please provide a numerical value in the range 1-30")
		exitCode = 1
		return
	}
	//if the input is a single digit, append a zero before it
	if dayInt < 10 {
		day = "0" + day
	}

	//dynamically call the method without using switch statements
	meth := reflect.ValueOf(solution).FieldByName(fmt.Sprintf("Day%s", day)).MethodByName("Run")
	meth.Call(nil)

}
