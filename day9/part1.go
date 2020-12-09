package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	invalidNumber := 0
	for i, v := range numbers {
		if i < 25 {
			continue
		}
		if !validateNumber(v, numbers[i-25:i]) {
			invalidNumber = v
		}
	}
	fmt.Println(invalidNumber)
}

func validateNumber(number int, batch []int) bool {
	valid := false
outside:
	for i, v := range batch {
		if valid {
			break
		}
		for ii := i + 1; ii < len(batch); ii++ {
			if (v + batch[ii]) == number {
				valid = true
				break outside
			}
		}
	}
	return valid
}
