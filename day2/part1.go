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

func splitInput(input string) (min, max int, letter byte, pass string) {
	splits := strings.Split(input, ": ")
	left := splits[0]
	pass = splits[1]

	splits2 := strings.Split(left, "-")
	var err error
	if min, err = strconv.Atoi(splits2[0]); err != nil {
		log.Fatal(err)
	}
	right := splits2[1]

	splits3 := strings.Split(right, " ")
	if max, err = strconv.Atoi(splits3[0]); err != nil {
		log.Fatal(err)
	}
	letter = splits3[1][0]
	// fmt.Println(min, max, letter, pass)
	return
}

func validatePassword(min, max int, letter byte, pass string) bool {
	count := 0
	for i := 0; i < len(pass); i++ {
		if letter == pass[i] {
			count++
		}
	}
	return count >= min && count <= max
}
