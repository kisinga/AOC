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
	// the sorted list becomes the longest variotion of arrangement
	sort.Ints(numbers)
	variations := findVariation(numbers)
	fmt.Println(variations)
}
func findVariation(input []int) uint64 {
	sort.Ints(input)
	// add the first and last vaue to the array
	input = append([]int{0}, append(input, input[len(input)-1]+3)...)
	m := map[int]uint64{}
	var next func(int) uint64
	next = func(idx int) uint64 {
		if m[idx] != 0 {
			return m[idx]
		}
		if idx == len(input)-1 {
			return 1
		}
		var sum uint64
		for i := idx + 1; i < len(input); i++ {
			if input[i]-input[idx] > 3 {
				break
			}
			n := next(i)
			sum += n
			m[i] = n
		}
		return sum
	}
	x := next(0)
	return x
}
