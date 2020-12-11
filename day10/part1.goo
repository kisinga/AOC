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

	sort.Ints(numbers)

	joltDifferenceCount := map[int]int{}
	for i, _ := range numbers {
		if i+1 >= len(numbers) {
			joltDifferenceCount[3]++
			continue
		}
		if i == 0 {
			joltDifferenceCount[numbers[0]-0]++
		}
		joltDifferenceCount[joltDifference(i, numbers)]++
	}
	fmt.Println(joltDifferenceCount[1] * joltDifferenceCount[3])
}
func joltDifference(pos int, joltArray []int) int {
	return joltArray[pos+1] - joltArray[pos]
}
