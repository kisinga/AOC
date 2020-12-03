package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	validCount := 0
	for scanner.Scan() {
		if validatePassword(splitInput(scanner.Text())) {
			validCount++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(validCount)

}

func splitInput(input string) (first, last int, letter byte, pass string) {
	splits := strings.Split(input, ": ")
	left := splits[0]
	pass = splits[1]

	splits2 := strings.Split(left, "-")
	var err error
	if first, err = strconv.Atoi(splits2[0]); err != nil {
		log.Fatal(err)
	}
	right := splits2[1]

	splits3 := strings.Split(right, " ")
	if last, err = strconv.Atoi(splits3[0]); err != nil {
		log.Fatal(err)
	}
	letter = splits3[1][0]
	// fmt.Println(min, max, letter, pass)
	return
}

func validatePassword(first, last int, letter byte, pass string) bool {
	if (pass[first-1] == letter && pass[last-1] != letter) || (pass[first-1] != letter && pass[last-1] == letter) {
		return true
	}

	return false
}
