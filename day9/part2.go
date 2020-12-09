package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	sum := 0
	for i, _ := range numbers {
		if i < 25 {
			continue
		}
		if set := findContiguousSet(invalidNumber, numbers[i-25:i]); len(set) > 0 {
			sort.Ints(set)
			sum = set[0] + set[len(set)-1]
		}

	}
	fmt.Println(sum)
}

func findContiguousSet(number int, batch []int) []int {
	for i, v := range batch {
		for ii := i + 1; ii < len(batch); ii++ {
			for x := i; x <= ii; x++ {
				sum := 0
				set := batch[i:x]
				for _, vv := range set {
					sum += vv
				}
				if sum == number {
					set = append(set, v)
					return set
				}
			}
		}
	}
	return nil
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
