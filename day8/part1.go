package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation int

const (
	acc operation = iota
	jmp operation = iota
	nop operation = iota
)

func (o operation) String() string {
	return string(o)
}

type instruction struct {
	operation
	value int
}
type line struct {
	executed bool
	instruction
}

var accumulator int

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	// make a map so that it's easy to search, just in case
	lines := []line{}
	for scanner.Scan() {
		lines = append(lines, extractLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	run(0, &lines[0], lines)

	fmt.Println(accumulator)
}
func extractLine(l string) line {
	parts := strings.Fields(l)
	val, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	op, err := parseOperation(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	return line{
		executed: false,
		instruction: instruction{
			operation: op,
			value:     val,
		},
	}
}
func run(index int, l *line, lines []line) {
	newIndex, err := executeLine(index, l)
	if err!= nil{
		return
	}
	run(newIndex, &lines[newIndex], lines)
}

func parseOperation(op string) (operation, error) {
	switch op {
	case "acc":
		return acc, nil
	case "jmp":
		return jmp, nil
	case "nop":
		return nop, nil
	default:
		return 0, errors.New("invalid string provided")
	}
}

func executeLine(index int, l *line) (int, error) {
	if l.executed {
		return 0, errors.New("line already executed")
	}
	return executeInstruction(index, l), nil
}

func executeInstruction(index int, l *line) int {
	l.executed = true
	switch l.operation {
	case acc:
		accumulator += l.value
		return index + 1
	case jmp:
		return index + l.value
	default:
		return index + 1
	}
}
